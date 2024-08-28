<template>
  <div id="personal-info">
    <h1 id="name">Simon Dahrs</h1>
    <h4>DevOps/Software engineer and information technology enthousiast</h4>
    <br />
    <div class="personal-links">
      <span class="link"
        ><a target="_blank" href="https://www.google.com/maps/place/Amsterdam"
          >📍 Amsterdam</a
        ></span
      >
      <span class="link"
        ><a target="_blank" href="https://www.linkedin.com/in/SimonDahrs"
          ><img class="icon" src="@/assets/img/linkedin.png" /></a
      ></span>
      <span class="link"
        ><a target="_blank" href="https://github.com/SimonDahrs"
          ><img class="icon" src="@/assets/img/github-mark.svg" /></a
      ></span>
    </div>
    <br />
  </div>
</template>
<style>
.personal-links {
  display: flex;
  column-gap: 1.5rem;
}

.icon {
  height: 1.6rem;
}

#personal-info {
  padding: 1.2em;
  background:
    radial-gradient(circle at 100% 100%, #ffffff 0, #ffffff 3px, transparent 3px) 0% 0%/8px 8px
      no-repeat,
    radial-gradient(circle at 0 100%, #ffffff 0, #ffffff 3px, transparent 3px) 100% 0%/8px 8px
      no-repeat,
    radial-gradient(circle at 100% 0, #ffffff 0, #ffffff 3px, transparent 3px) 0% 100%/8px 8px
      no-repeat,
    radial-gradient(circle at 0 0, #ffffff 0, #ffffff 3px, transparent 3px) 100% 100%/8px 8px
      no-repeat,
    linear-gradient(#ffffff, #ffffff) 50% 50% / calc(100% - 10px) calc(100% - 16px) no-repeat,
    linear-gradient(#ffffff, #ffffff) 50% 50% / calc(100% - 16px) calc(100% - 10px) no-repeat,
    linear-gradient(145deg, transparent 50%, #dce5f7 100%);
  border-radius: 8px;
}
</style>
