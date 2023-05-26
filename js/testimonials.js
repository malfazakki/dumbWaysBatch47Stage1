// class Testimonial {
//   #quote = "";
//   #image = "";

//   constructor(quote, image) {
//     this.#quote = quote;
//     this.#image = image;
//   }

//   get quote() {
//     return this.#quote;
//   }

//   get image() {
//     return this.#image;
//   }

//   // This is an abstract method that subclasses will implement
//   get author() {
//     throw new Error("getAuthor() method must be implemented");
//   }

//   // This is a polymorphic method that can take any subclasses of Testimonial
//   get testimonialHTML() {
//     return `<div class="card">
//                 <img src="${this.image}" />
//                 <p class="quote">${this.quote}</p>
//                 <p class="author">- ${this.author}</p>
//             </div>
//         `;
//   }
// }

// // Subclass
// class AuthorTestimonials extends Testimonial {
//   #author = "";

//   constructor(author, quote, image) {
//     super(quote, image);
//     this.#author = author;
//   }

//   get author() {
//     return this.#author;
//   }
// }

// // Subclass
// class CompanyTestimonials extends Testimonial {
//   #company = "";

//   constructor(company, quote, image) {
//     super(quote, image);
//     this.#company = company;
//   }

//   get author() {
//     return this.#company + " Company";
//   }
// }

// const testimonial1 = new AuthorTestimonials(
//   "Malfazakki",
//   "Jagalah Kebersihan",
//   "assets/images/profile/maxresdefault.jpg"
// );
// const testimonial2 = new AuthorTestimonials(
//   "Cintara Surya",
//   "Keren cuys!!",
//   "assets/images/profile/maxresdefault.jpg"
// );
// const testimonial3 = new CompanyTestimonials(
//   "Yamaha",
//   "Desain Yang Mantap!! 🔥🔥🔥",
//   "assets/images/profile/maxresdefault.jpg"
// );

// let testimonialData = [testimonial1, testimonial2, testimonial3];
// let testimonialHTML = "";

// for (let i = 0; i < testimonialData.length; i++) {
//   testimonialHTML += testimonialData[i].testimonialHTML;
// }

// document.getElementById("card").innerHTML = testimonialHTML;

// Array of Object untuk menampung Data Testimoni start
const testimonialData = [
  {
    author: "Malfazakki",
    quote: "Tetaplah menjadi baik",
    image: "assets/images/profile/maxresdefault.jpg",
    rating: 5,
  },
  {
    author: "Alfonso Guefarra",
    quote: "Kerreeeeennnnn ❤️❤️",
    image: "assets/images/profile/1.jpg",
    rating: 4,
  },
  {
    author: "Masbro Gans",
    quote: "Sangat memuaskan dan sesuai dengan jangka waktu yang diberikan",
    image: "assets/images/profile/2.jpg",
    rating: 4,
  },
];
// Array of Object untuk menampung Data Testimoni start

// Function untuk Looping data Testimoni start
function allTestimonials() {
  let testimonialHTML = "";

  testimonialData.forEach(function (item) {
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
// Function untuk Looping data Testimoni start

// menjalankan function testimoni
allTestimonials();

// function untuk filter testimoni start
function filterTestimonials(rating) {
  let testimonialsHTML = "";

  const testimonialFiltered = testimonialData.filter(function (item) {
    return item.rating === rating;
  });

  if (testimonialFiltered.length === 0) {
    testimonialsHTML += `<h1>Data not Found!</h1>`;
  } else {
    testimonialFiltered.forEach(function (item) {
      testimonialsHTML += `<div class="card">
                            <img src="${item.image}" />
                            <p class="quote">${item.quote}</p>
                            <p class="author">- ${item.author}</p>
                            <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
                          </div>
      `;
    });
  }
  document.getElementById("card").innerHTML = testimonialsHTML;
}
