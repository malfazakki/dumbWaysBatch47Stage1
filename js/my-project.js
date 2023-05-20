let dataProject = [];

function addProject(event) {
  event.preventDefault();

  let title = document.getElementById("project-name").value;
  let startDate = document.getElementById("start-date").value;
  let endDate = document.getElementById("end-date").value;
  let description = document.getElementById("description").value;
  let image = document.getElementById("image-upload").files;

  if (
    title === "" ||
    startDate === "" ||
    endDate === "" ||
    description === "" ||
    image === ""
  ) {
    alert("Pastikan semua kolom formulir terisi!");
  }

  image = URL.createObjectURL(image[0]);
  console.log(image);

  let project = {
    title,
    startDate,
    endDate,
    description,
    image,
  };

  dataProject.push(project);
  console.log(dataProject);

  renderProject();
}

function renderProject() {
  for (let index = 0; index < dataProject.length; index++) {
    document.getElementById("contents").innerHTML += `
      <div class="card">
        <img src="${dataProject[index].image}" alt="${dataProject[index].title}" />
        <h3 class="judul-artikel">
          <a href="project-detail.html" target="_blank">${dataProject[index].title}</a>
        </h3>
        <div class="durasi">
          <p class="durasi">${dataProject[index].startDate}</p>
          <p class="durasi">${dataProject[index].endDate}</p>
        </div>
        <p class="deskripsi-artikel">
          ${dataProject[index].description}
        </p>
        <div class="techimages">
          ${nodeJs}
        </div>
        <div class="btn-artikel">
          <button class="btn-edit">edit</button>
          <button class="btn-delete">delete</button>
        </div>
      </div>
    `;
  }
}
