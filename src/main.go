//フロントへ返す値のキー
//locationid
//bussstateMsg
//remaintimeMsg
//timetableMsg
//lastbusMsg

package main

import (
  "fmt"
  "log"
  "io"
  "time"
  "os"
  "math"
  "encoding/csv"
  "net/http"
  "html/template"
  "strconv"
  "strings"
)

type Bus struct {
  id string
  latitude float64
  longitude float64
}
var buses map[string]*Bus = map[string]*Bus{} //struct Bus map


//targetがbegin~end間にあるか判定する
func inTime(begin time.Time, end time.Time, target time.Time) bool {
  return begin.Before(target) && end.After(target)
}
//string"x:xx~x:xx" -> [x:xx(Time),x:xx(Time)]
func timeParse(str string) []time.Time {
  result := []time.Time{}
  t := time.Now()
  t = time.Date(2018, 8, 3, 10, 41, 23, 0, time.Local)
  year := t.Year() //現在年
  month := t.Month() //現在月
  day := t.Day() //現在日
  strs := strings.Split(str,"~")
  for i:=0; i<len(strs); i++ {
    timestr := strings.Split(strs[i],":")
    hour, _ := strconv.Atoi(timestr[0])
    min, _ := strconv.Atoi(timestr[1])
    timedata := time.Date(year,month,day,hour,min,0,0,time.Local)
    result = append(result,timedata)
  }
  return result
}
//バスの緯度経度からバスの位置を推定
func calcBusLocation(bus *Bus, csvPath string) string{
  if bus==nil {
    return ""
  }
  var approx float64 = math.MaxFloat64
  var locationid int

  file, err := os.Open(csvPath)
  failOnError(err)
  defer file.Close()

  reader := csv.NewReader(file)
  reader.Read()
  for i:=1;;i++{
    record, err := reader.Read()
    if err == io.EOF {
      break
    } else if err != nil {
      panic(err)
    }
    lat, _ := strconv.ParseFloat(record[1],64)
    lon, _ := strconv.ParseFloat(record[2],64)
    x := bus.longitude
    y := bus.latitude
    dist := math.Sqrt((lon-x)*(lon-x)+(lat-y)*(lat-y))
    if(approx>dist) {
      approx = dist
      locationid = i
    }
  }
  return "location"+strconv.Itoa(locationid)
}


//Bus情報更新ページHandler
func setPosition(w http.ResponseWriter ,r *http.Request){
  r.ParseForm()
  var id string = r.Form["id"][0]
  var lat, _ = strconv.ParseFloat(r.Form["latitude"][0],64)
  var lon, _ = strconv.ParseFloat(r.Form["longitude"][0],64)
  if _, exist := buses[id]; exist{
    //このバスがすでに登録されていたら
    buses[id].latitude = lat
    buses[id].longitude = lon
  } else {
    //バスが未登録であれば
    var bus Bus = Bus{id,lat,lon}
    buses[id] = &bus
  }
}

//バス情報確認ページHandler
func getPosition(w http.ResponseWriter ,r *http.Request){
  fmt.Fprintf(w, `
      <!DOCTYPE html>
      <html>
      <body>
  `)
  for _, bus := range buses {
    fmt.Fprintf(w, `
      <p> id:%s </p>
      <p> latitude:%f </p>
      <p> longitude:%f </p>
      <br>
    `,bus.id,bus.latitude,bus.longitude)
  }
  fmt.Fprintf(w, `
      </body>
      </html>
  `)
}


func topPageHandler(w http.ResponseWriter ,r *http.Request){
  r.ParseForm()
  var data map[string]string = map[string]string{} //クライアントへの返却データ群
  var stopid string //ユーザが指定したバス停id

  if _, exist := r.Form["stopid"]; exist {
    //stopidが送信されれば
    stopid = r.Form["stopid"][0] //stop? (string)
    t := time.Now() //現在時刻
    t = time.Date(2018, 8, 3, 10, 41, 23, 0, time.Local)
    week := t.Weekday().String() //現在曜日(string)
    var timetableCSV string //運行時間のcsv
    var locationsCSV string //gps測定地点のcsv

    switch week {
    case "Sunday":
      timetableCSV = ""
      break
    case "Saturday":
      timetableCSV = "../webContents/data/bustimesat.csv"
      break
    default:
      timetableCSV = "../webContents/data/bustime.csv"
    }
    locationsCSV = "../webContents/data/locations.csv"

    if timetableCSV != ""{
      file, err := os.Open(timetableCSV)
      failOnError(err)
      defer file.Close()
      reader := csv.NewReader(file)
      var id int
      switch stopid {
      case "stop1":
        id = 1
        break
      case "stop2":
        id = 2
        break
      case "stop3":
        id = 3
      case "stop4":
        id = 4
      default:
        data["busstateMsg"] = ""
        data["remaintimeMsg"] = ""
        data["timetableMsg"] = ""
        data["lastbusMsg"] = ""
        //template送信処理
        t := template.Must(template.ParseFiles("../webContents/index.tmpl"))
        if err := t.ExecuteTemplate(w,"index.tmpl",data); err != nil {
          log.Fatal(err)
        }
      }
      for i:=0; i<id-1; i++ {
        reader.Read()
      }
      record, err := reader.Read()
      failOnError(err)

      now := t
      var text string = ""
      var isInService bool = false
      var remainTime time.Duration
      var isSetRemainTime = false
      for i:=1; i<len(record)-1; i++ {
        text = text + record[i]
        times := timeParse(record[i])
        if len(times)==2 && inTime(times[0],times[1],now) {
          isInService = true
        } else {
          if now.Before(times[0]) && !isSetRemainTime {
            remainTime = times[0].Sub(now)
            isSetRemainTime = true
          }
        }
        if i+1 < len(record)-1 {
          text = text + "\n"
        }
      }

      //送信データセット
      for _,bus := range buses {
        data[calcBusLocation(bus,locationsCSV)] = "t"
        //data["location1"] = "t"
      }
      switch isInService {
      case true:
        data["busstateMsg"] = "現在、運行中です。"
        data["remaintimeMsg"] = ""
      default:
        data["busstateMsg"] = "現在、運行していません。"
        if !isSetRemainTime {
          data["remaintimeMsg"] = "本日の運行は終了しました。"
        }else {
          remainHour := strconv.Itoa(int(remainTime.Hours()))
          remainMin := strconv.Itoa(int(remainTime.Minutes())%60)
          data["remaintimeMsg"] = "あと"+remainHour+"時間"+remainMin+"分で運転を再開します。"
        }
      }
      data["timetableMsg"] = text
      data["lastbusMsg"] = record[len(record)-1]

    } else {
      //日曜日
      data["locationid"] = ""
      data["busstateMsg"] = "本日は、運行していません。"
      data["remaintimeMsg"] = ""
      data["timetableMsg"] = ""
      data["lastbusMsg"] = ""
    }
  }

  //template送信処理
  t := template.Must(template.ParseFiles("../webContents/index.tmpl"))
  if err := t.ExecuteTemplate(w,"index.tmpl",data); err != nil {
    log.Fatal(err)
  }
}

func failOnError(err error){
  if err != nil {
    log.Fatal("Error",err)
  }
}

func main(){
  //buses["test"] = &Bus{"test",1.1,1.1}
  http.HandleFunc("/setPosition", setPosition)
  http.HandleFunc("/getPosition", getPosition)
  http.HandleFunc("/", topPageHandler)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../webContents"))))
  log.Fatal(http.ListenAndServe(":8080", nil))
}
