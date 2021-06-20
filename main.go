package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
    port := os.Getenv("envport")
	if port == "" {
		port = "8888"
	log.Printf("Docker web running on port %s", port)
	} else {
	log.Printf("Docker web running on port %s", port)
	}

	http.HandleFunc("/", index)
	http.ListenAndServe(fmt.Sprintf(`:%s`, port), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	title := os.Getenv("envtitle")
	if title == ""{
		title = "Shamim"
	}
	msg := os.Getenv("envmsg")
	if msg ==""{
		msg= "just for fun.This site is running my home Laptop"
	}

	hundred := `100%`
	fifty := `50%`
	forty := `40%`

  hName, err := os.Hostname()
  if err != nil {
    log.Printf(" Hostname error %s", err )
  }
  userAgent := r.UserAgent()
  userIp := r.RemoteAddr


	log.Printf("req from %s\n", r.Host)

    w.Header().Set("Content-Type", "text/html")
	
	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<style>
body, html {
  height: %s;
  margin: 0;
}

.bgblack {
    background-color: #00000085;
    position: absolute;
    top: 0;
    left: 0;
    width: %s;
    height: %s;
}

.bgimg {
  background-image: url('https://i.pinimg.com/originals/77/ea/6d/77ea6d59c910b1fcd4cc3a32f52c5189.jpg');
  height: %s;
  background-position: center;
  background-size: cover;
  position: relative;
  color: white;
  font-family: "Courier New", Courier, monospace;
  font-size: 25px;
}

.topleft {
  position: absolute;
  top: 50px;
  left: 35px;
  width: 240px;
  font-size: 14px;
  overflow-y: scroll;
  height: 150px;
}

/*
 * scrollbar STYLE 1
 */

#scrollbar::-webkit-scrollbar-track
{
  -webkit-box-shadow: inset 0 0 6px rgba(0,0,0,0.3);
  border-radius: 10px;
  // background-color: #F5F5F5;
}

#scrollbar::-webkit-scrollbar
{
  width: 12px;
  // background-color: #F5F5F5;
}

#scrollbar::-webkit-scrollbar-thumb
{
  border-radius: 10px;
  -webkit-box-shadow: inset 0 0 6px rgba(0,0,0,.3);
  background-color: #555;
}

.bottomright {
  position: absolute;
  bottom: 0;
  right: 16px;
}

.middle {
  position: absolute;
  top: %s;
  left: %s;
  transform: translate(-%s, -%s);
  text-align: center;
}

hr {
  margin: auto;
  width: %s;
}
</style>
<body>

<div class="bgimg">
<div class="bgblack"></div>
  <div id="scrollbar" class="topleft">
    <p>%s</p>
    <p>"Hostname %s"</p>
    <p>"Server Running on %s"</p>
    <p>"User-Agent from %s"</p>
    <p>"User-IP from %s"</p>
  </div>
  <div class="middle">
    <h1>WOW I'm Working</h1>
    <hr>
    <p>Your message : %s</p>
  </div>
  <div class="bottomright">
    <p>Thank you for believing in me ( Shamim )</p>
  </div>
</div>
<script>
var objDiv = document.getElementById("scrollbar");
objDiv.scrollTop = objDiv.scrollHeight;
</script>
</body>
</html>`, hundred , hundred, hundred , hundred, fifty, fifty, fifty, fifty, forty ,title ,hName, r.Host, userAgent ,userIp , msg)
}

