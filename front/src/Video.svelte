<script>
	export let videoUrl;
	export let videoName;
</script>

<style>

video {
   width: auto !important;
   height: 100px !important;
}

.video {
  height: 100px;
  padding: 5px
  /*margin-bottom:5px;*/
}

video.hover-shadow {
  transition: 0.3s;
}

.hover-shadow:hover {
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
}

.cursor {
  cursor: pointer;
}
</style>

<div class="video">
 <video controls class="hover-shadow cursor">
  <source src={videoUrl}>
  </video> 
</div>

 
