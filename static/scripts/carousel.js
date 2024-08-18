let slideIndex = 0;

function showSlides() {
    const slides = document.querySelectorAll(".carousel-slide img");
    if (slides.length > 0) {
        for (let i = 0; i < slides.length; i++) {
            slides[i].style.display = "none";
        }
        slideIndex++;
        if (slideIndex > slides.length) {
            slideIndex = 1;
        }
        slides[slideIndex - 1].style.display = "block";
    }
}

function moveSlide(n) {
    const slides = document.querySelectorAll(".carousel-slide img");
    if (slides.length > 0) {
        slideIndex += n;
        if (slideIndex > slides.length) {
            slideIndex = 1;
        }
        if (slideIndex < 1) {
            slideIndex = slides.length;
        }
        showSlides();
    }
}

document.addEventListener("DOMContentLoaded", function() {
    showSlides();
    setInterval(() => moveSlide(1), 5000); // Change slide every 5 seconds
});
