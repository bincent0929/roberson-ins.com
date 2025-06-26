/*
    this complains about thing being redeclared but it's only because
    it's trying to read the js from the compiled file
*/
let slideIndex: number = 1;
showSlides(slideIndex)

function plusSlides(n: number): void {
    slideIndex += n;
    showSlides(slideIndex);
}

function showSlides(n: number): void {
  const slides = document.getElementsByClassName("mySlides") as HTMLCollectionOf<HTMLElement>;

  if (n > slides.length) {
    slideIndex = 1;
  }
  if (n < 1) {
    slideIndex = slides.length;
  }

  for (let i = 0; i < slides.length; i++) {
    slides[i].style.display = "none";
  }

  slides[slideIndex - 1].style.display = "block";
}