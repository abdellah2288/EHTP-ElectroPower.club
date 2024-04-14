const slideContainers = document.querySelectorAll(".slide-container");
slideContainers.forEach(slideContainer => {
  const buttonNext = slideContainer.querySelector(".slide-button#next");
  const buttonPrev = slideContainer.querySelector(".slide-button#prev");
  const slides = slideContainer.querySelectorAll(".img-container")
  console.log(slides.length)
  buttonNext.addEventListener("click",()=>{slideForward(slides);});
  buttonPrev.addEventListener("click",()=>{slideBackward(slides)});
  setInterval(() => {
    slideForward(slides);
  }, 10000);
});
function slideForward(slides)
{
    for(index = 0;index<slides.length;index++)
    {
      if(typeof slides[index].dataset.active !== "undefined") 
      {
        nextIndex = index + 1;
        if(nextIndex<0) nextIndex = slides.length -1;
        if(nextIndex>=slides.length) nextIndex = 0;
        slides[nextIndex].dataset.active=true;
        delete slides[index].dataset.active;
        break;
      }
    }
}
function slideBackward(slides)
{
    for(index = 0;index<slides.length;index++)
    {
      if(typeof slides[index].dataset.active !== "undefined") 
      {
        nextIndex = index - 1;
        if(nextIndex<0) nextIndex = slides.length -1;
        if(nextIndex>=slides.length) nextIndex = 0;
        slides[nextIndex].dataset.active=true;
        delete slides[index].dataset.active;
        break;
      }
    }
}
