<template>
  <div id="personal-info">
    <h1 id="name">Simon Dahrs</h1>
    <h4>DevOps/Software engineer passionate about all things tech.</h4>
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
          ><img class="icon" src="@/assets/img/github-mark-white.svg" /></a
      ></span>
    </div>
    <br />
  </div>
</template>
<style>
.personal-links {
  display: flex;
  column-gap: 1.5rem;
  width: 100%;
  align-items: center;
  flex-wrap: wrap;
}

.icon {
  max-height: 1.6rem;
}

#personal-info {
  padding: 1.2em;
}

@media only screen and (max-width: 768px) {
  /* .link:first-child {
    width: 100%;
    margin-bottom: 5px;
  } */
  #personal-info {
    background-image: linear-gradient(178deg, var(--vt-c-background-left) 90%, var(--vt-c-background-right) 90%);
  }
}
</style>
