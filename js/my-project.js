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
    return alert("Pastikan semua kolom formulir terisi!");
  }

  nodeIcon = '<i class="fa-brands fa-node-js fa-xl"></i>';

  let nodeChecked = document.getElementById("nodeJs").checked ? nodeIcon : "";
  let reactChecked = document.getElementById("react").checked
    ? '<i class="fa-brands fa-react fa-xl"></i>'
    : "";
  let bootstrapChecked = document.getElementById("bootstrap").checked
    ? '<i class="fa-brands fa-bootstrap fa-xl"></i>'
    : "";
  let laravelChecked = document.getElementById("laravel").checked
    ? '<i class="fa-brands fa-laravel fa-xl"></i>'
    : "";

  image = URL.createObjectURL(image[0]);
  console.log(image);

  let project = {
    title,
    startDate,
    endDate,
    description,
    image,
    nodeChecked,
    reactChecked,
    bootstrapChecked,
    laravelChecked,
  };

  dataProject.push(project);
  console.log(dataProject);

  renderProject();
}

function renderProject() {
  document.getElementById("contents").innerHTML = "";

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
          ${dataProject[index].nodeChecked}
          ${dataProject[index].reactChecked}
          ${dataProject[index].bootstrapChecked}
          ${dataProject[index].laravelChecked}
        </div>
        <div class="btn-artikel">
          <button class="btn-edit">edit</button>
          <button class="btn-delete">delete</button>
        </div>
      </div>
    `;
  }
}
