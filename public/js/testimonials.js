const promise = new Promise((resolve, reject) => {
  const xhr = new XMLHttpRequest();
  xhr.open("GET", "https://api.npoint.io/7d80527f33dd94fe8f24", true);
  xhr.onload = () => {
    if (xhr.status === 200) {
      resolve(JSON.parse(xhr.response));
    } else {
      reject("Error loading data.");
    }
  };
  xhr.onerror = () => {
    reject("Network Error");
  };
  xhr.send();
});

async function getAllTestimonials() {
  const response = await promise;

  let testimonialHTML = "";
  response.forEach(function (item) {
    testimonialHTML += `<div class="card col card-project px-0">
            <img src="${item.image}" class="card-img-top" alt="..."
              style="height: 200px; object-fit: cover; width:100%;">
            <div class="card-body">
              <p class="quote"><i>${item.quote}</i></p>
              <p class="author text-end mb-0">${item.author}</p>
              <p class="author text-end mb-2">${item.rating} <i class="fa-solid fa-star"></i></p>
            </div>
          </div>
      `;
  });
  document.getElementById("card").innerHTML = testimonialHTML;
}

getAllTestimonials();

async function getFilteredTestimonials(rating) {
  const response = await promise;

  const testimonialFiltered = response.filter((item) => {
    return item.rating === rating;
  });

  console.log(testimonialFiltered);

  let testimonialHTML = "";

  if (testimonialFiltered.length === 0) {
    testimonialHTML = `<div class="w-100"><p>Data not found!</p></div>`;
  } else {
    testimonialFiltered.forEach((item) => {
      testimonialHTML += `<div class="card col card-project px-0">
            <img src="${item.image}" class="card-img-top" alt="..."
              style="height: 200px; object-fit: cover; width:100%;">
            <div class="card-body">
              <p class="quote"><i>${item.quote}</i></p>
              <p class="author text-end mb-0">${item.author}</p>
              <p class="author text-end mb-2">${item.rating} <i class="fa-solid fa-star"></i></p>
            </div>
          </div>
      `;
    });
  }

  document.getElementById("card").innerHTML = testimonialHTML;
}
