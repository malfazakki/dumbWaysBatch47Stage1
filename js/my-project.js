let dataProject = [];

function addProject(event) {
  event.preventDefault();

  let title = document.getElementById("project-name").value;
  let startDate = document.getElementById("start-date").value;
  let endDate = document.getElementById("end-date").value;
  let description = document.getElementById("description").value;
  let image = document.getElementById("image-upload").files;
  let imageCheck = document.getElementById("image-upload").value;

  if (
    title === "" ||
    startDate === "" ||
    endDate === "" ||
    description === "" ||
    imageCheck == ""
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
  const selesai = moment(startDate);
  const mulai = moment(endDate);
  const dayDuration = mulai.diff(selesai, "days");
  const monthDuration = mulai.diff(selesai, "months");
  const yearsDuration = mulai.diff(selesai, "years");
  let duration;

  if (monthDuration > 0) {
    duration = `duration: ${monthDuration} months`;
  } else if (yearsDuration > 0) {
    duration = `duration: ${yearsDuration} years`;
  } else {
    duration = `duration: $ {dayDuration} days`;
  }
  console.log(image);

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
}

function renderProject() {
  document.getElementById("contents").innerHTML = "";

  for (let index = 0; index < dataProject.length; index++) {
    document.getElementById("contents").innerHTML += `
      <div class="card">
        <img src="${dataProject[index].image}" alt="${
      dataProject[index].title
    }" />
        <h3 class="judul-artikel">
          <a href="project-detail.html" target="_blank">${
            dataProject[index].title
          }</a>
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
  let monthName = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
  ];
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

// function getDistanceTime(time) {
//   let timeNow = new Date();
//   let timePost = time;

//   // waktu sekarang - waktu post
//   let distance = timeNow - timePost; // hasilnya milidetik
//   console.log(distance);

//   let milisecond = 1000; // milisecond
//   let secondInHours = 3600; // 1 jam 3600 detik
//   let hoursInDays = 24; // 1 hari 24 jam
//   let daysInWeek = 7; // 1 minggu 7 hari
//   let weeksInMonth = 4; // 1 bulan 4 minggu
//   let monthsInYear = 12; // 1 tahun 12 bulan

//   let distanceDay = Math.floor(
//     distance / (milisecond * secondInHours * hoursInDays)
//   );
//   let distanceWeek = Math.floor(
//     distance / (milisecond * secondInHours * hoursInDays * daysInWeek)
//   );
//   let distanceMonth = Math.floor(
//     distance /
//       (milisecond * secondInHours * hoursInDays * daysInWeek * weeksInMonth)
//   );
//   let distanceYear = Math.floor(
//     distance /
//       (milisecond *
//         secondInHours *
//         hoursInDays *
//         daysInWeek *
//         weeksInMonth *
//         monthsInYear)
//   );

//   if (distanceYear > 0) {
//     return `duration: ${distanceYear} years`;
//   } else if (distanceMonth > 0) {
//     return `duration: ${distanceMonth} months`;
//   } else if (distanceWeek > 0) {
//     return `duration: ${distanceWeek} weeks`;
//   } else if (distanceDay > 0) {
//     return `duration: ${distanceDay} days`;
//   } else {
//     return alert("Invalid Date input");
//   }
// }
