let dataProject = [];

function addProject(event) {
  event.preventDefault();

  let title = document.getElementById("project-name").value;
  let startDate = document.getElementById("start-date").value;
  let endDate = document.getElementById("end-date").value;
  let description = document.getElementById("description").value;
  let image = document.getElementById("image-upload").files;
  let imageCheck = document.getElementById("image-upload").value;

  if (title === "" || startDate === "" || endDate === "" || description === "" || imageCheck == "") {
    return alert("Pastikan semua kolom formulir terisi!");
  }

  // Untuk Icon Techstack
  let nodeChecked = document.getElementById("nodeJs").checked ? '<i class="fa-brands fa-node-js fa-xl"></i>' : "";
  let reactChecked = document.getElementById("react").checked ? '<i class="fa-brands fa-react fa-xl"></i>' : "";
  let bootstrapChecked = document.getElementById("bootstrap").checked
    ? '<i class="fa-brands fa-bootstrap fa-xl"></i>'
    : "";
  let laravelChecked = document.getElementById("laravel").checked ? '<i class="fa-brands fa-laravel fa-xl"></i>' : "";

  // membuat url dari image
  image = URL.createObjectURL(image[0]);

  // Menghitung Durasi dengan MomentJs
  const mulai = moment(startDate);
  const selesai = moment(endDate);
  const daysDuration = selesai.diff(mulai, "days");
  const monthsDuration = selesai.diff(mulai, "months");
  const weeksDuration = selesai.diff(mulai, "weeks");
  const yearsDuration = selesai.diff(mulai, "years");
  let duration;

  // menghitung durasi sesuai input date
  if (yearsDuration > 0) {
    duration = `duration: ${yearsDuration} years`;
  } else if (monthsDuration > 0) {
    duration = `duration: ${monthsDuration} months`;
  } else if (weeksDuration > 0) {
    duration = `duration: ${weeksDuration} weeks`;
  } else if (daysDuration > 0) {
    duration = `duration: ${daysDuration} days`;
  } else {
    return alert("invalid date input!");
  }

  let project = {
    title,
    duration,
    description,
    image,
    nodeChecked,
    reactChecked,
    bootstrapChecked,
    laravelChecked,
    postAt: new Date(),
    author: "Malfazakki",
  };

  dataProject.push(project);
  console.log(dataProject);

  renderProject();
  resetForm();
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
        <div class="duration">
          <p>${dataProject[index].duration}</p>
          <p>Posted On:  ${getFullTime(dataProject[index].postAt)}</p>
          <p>by: ${dataProject[index].author}</p>
          <div>
         
          </div>
        </div>
        <div class="deskripsi-artikel">
          <p>
           ${dataProject[index].description}
          </p>
        </div>      
        <div class="tech-images">
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

function getFullTime(time) {
  let monthName = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];
  // console.log(monthName[8]);

  let date = time.getDate();
  // console.log(date);

  let monthIndex = time.getMonth();
  // console.log(monthIndex);

  let year = time.getFullYear();
  // console.log(year);

  let hours = time.getHours();
  let minutes = time.getMinutes();
  // console.log(minutes);

  if (hours <= 9) {
    hours = "0" + hours;
  } else if (minutes <= 9) {
    minutes = "0" + minutes;
  }

  return `${date} ${monthName[monthIndex]} ${year} ${hours}:${minutes} WIB`;
}

function resetForm() {
  document.getElementById("myForm").reset();
}
