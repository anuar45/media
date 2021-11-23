<script>
	import { onMount } from 'svelte';
	import { getMediaItems } from "./api/media";

	let mediaItemList = [1,2,3,5];
	let slideIndex = 1;

	onMount(async () => {
    const res = await getMediaItems("/");
		console.log(res)
    mediaItemList = res;
  });

// Open the Modal
function openModal() {
  document.getElementById("myModal").style.display = "block";
}

// Close the Modal
function closeModal() {
  document.getElementById("myModal").style.display = "none";
}

// showSlides(slideIndex);

// // Next/previous controls
// function plusSlides(n) {
//   showSlides(slideIndex += n);
// }

// // Thumbnail image controls
// function currentSlide(n) {
//   showSlides(slideIndex = n);
// }

// function showSlides(n) {
//   var i;
//   var slides = document.getElementsByClassName("mySlides");
//   var dots = document.getElementsByClassName("demo");
//   var captionText = document.getElementById("caption");
//   if (n > slides.length) {slideIndex = 1}
//   if (n < 1) {slideIndex = slides.length}
//   for (i = 0; i < slides.length; i++) {
//     slides[i].style.display = "none";
//   }
//   for (i = 0; i < dots.length; i++) {
//     dots[i].className = dots[i].className.replace(" active", "");
//   }
//   slides[slideIndex-1].style.display = "block";
//   dots[slideIndex-1].className += " active";
//   captionText.innerHTML = dots[slideIndex-1].alt;
// }

</script>


<main>
	<div id="app" class="container">
		<!-- Images used to open the lightbox -->
		<div class="row">
			{#each mediaItemList as mediaItem}
				<div class="column">
					<img src={mediaItem.path} alt={mediaItem.name} style="width:100%" on:click={openModal} class="hover-shadow cursor">
				</div>
			{/each}
		</div>

		<!-- The Modal/Lightbox -->
		<div id="myModal" class="modal">
			<span class="close cursor" on:click={closeModal}>&times;</span>
			<div class="modal-content">

				<!-- <div class="mySlides">
					{#each mediaItemList as mediaItem}
					<div class="numbertext">1/4</div>
					<img src={mediaItem.path} alt={mediaItem.name} style="width:100%">
					{/each}
				</div> -->

				<!-- Next/previous controls -->
				<!-- <a class="prev" on:click={plusSlides(-1)}>&#10094;</a>
				<a class="next" on:click={plusSlides(1)}>&#10095;</a>

				<!-- Caption text -->
				<!-- <div class="caption-container">
					<p id="caption"></p>
				</div> --> -->

				<!-- Thumbnail image controls -->
				<!-- <div class="column">
					<img class="demo" src="img1.jpg" onclick="currentSlide(1)" alt="Nature">
				</div> -->
				<!-- 
				<div class="column">
					<img class="demo" src="img2.jpg" onclick="currentSlide(2)" alt="Snow">
				</div> -->
			</div>
		</div> 
	</div>
</main>

<style>
main {
  font-family: Verdana, sans-serif;
  margin: 0;
}

* {
  box-sizing: border-box;
}

.row > .column {
  padding: 8px 8px;
}

.row:after {
  content: "";
  display: table;
  clear: both;
}

.column {
  float: left;
  width: 25%;
}

/* The Modal (background) */
.modal {
  display: none;
  position: fixed;
  z-index: 1;
  padding-top: 100px;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: black;
}

/* Modal Content */
.modal-content {
  position: relative;
  background-color: #fefefe;
  margin: auto;
  padding: 0;
  width: 90%;
  max-width: 1200px;
}

/* The Close Button */
.close {
  color: white;
  position: absolute;
  top: 10px;
  right: 25px;
  font-size: 35px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: #999;
  text-decoration: none;
  cursor: pointer;
}

.mySlides {
  display: none;
}

.cursor {
  cursor: pointer;
}

/* Next & previous buttons */
.prev,
.next {
  cursor: pointer;
  position: absolute;
  top: 50%;
  width: auto;
  padding: 16px;
  margin-top: -50px;
  color: white;
  font-weight: bold;
  font-size: 20px;
  transition: 0.6s ease;
  border-radius: 0 3px 3px 0;
  user-select: none;
  -webkit-user-select: none;
}

/* Position the "next button" to the right */
.next {
  right: 0;
  border-radius: 3px 0 0 3px;
}

/* On hover, add a black background color with a little bit see-through */
.prev:hover,
.next:hover {
  background-color: rgba(0, 0, 0, 0.8);
}

/* Number text (1/3 etc) */
.numbertext {
  color: #f2f2f2;
  font-size: 12px;
  padding: 8px 12px;
  position: absolute;
  top: 0;
}

img {
  margin-bottom: -4px;
}

.caption-container {
  text-align: center;
  background-color: black;
  padding: 2px 16px;
  color: white;
}

.demo {
  opacity: 0.6;
}

.active,
.demo:hover {
  opacity: 1;
}

img.hover-shadow {
  transition: 0.3s;
}

.hover-shadow:hover {
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}
</style>