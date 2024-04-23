const hackathonDate = new Date("Apr 24, 2024 18:00:00").getTime();
const countdownOutput = document.querySelector("p#countdownText");
function updateCountdown() {
var currentDate = new Date().getTime();

var distance = hackathonDate - currentDate;
var days = Math.floor(distance / (1000 * 60 * 60 * 24));
var hours = Math.floor((distance % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
var minutes = Math.floor((distance % (1000 * 60 * 60)) / (1000 * 60));
var seconds = Math.floor((distance % (1000 * 60)) / 1000);

countdownOutput.innerText = days + "d " + hours + "h "
            + minutes + "m " + seconds + "s "; 
 
}
setInterval(()=>{
  updateCountdown();
}
, 1000);

