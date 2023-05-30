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
    testimonialHTML += `<div class="card">
                          <img src="${item.image}" />
                          <p class="quote">${item.quote}</p>
                          <p class="author">- ${item.author}</p>
                          <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
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

  //   console.log(testimonialFiltered);

  let testimonialHTML = "";

  if (testimonialFiltered.length === 0) {
    testimonialHTML = "<h1>Data not found!</h1>";
  } else {
    testimonialFiltered.forEach((item) => {
      testimonialHTML += `<div class="card">
                                <img
                                    src="${item.image}"
                                    class="profile-testimonial"
                                />
                                <p class="quote">${item.quote}</p>
                                <p class="author">- ${item.author}</p>
                                <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
                            </div>
                        `;
    });
  }

  document.getElementById("card").innerHTML = testimonialHTML;
}
