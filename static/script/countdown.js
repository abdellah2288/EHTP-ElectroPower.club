const hackathonDate = new Date("2024-04-15");
const countdownOutput = document.querySelector("p#countdownText");
function updateCountdown() {
 let currentDate = new Date();

let distance = hackathonDate - currentDate;
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

