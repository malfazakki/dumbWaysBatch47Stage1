let dataBlog = [];

function addBlog(event) {
  event.preventDefault();

  let title = document.getElementById("project-name").value;
  let content = document.getElementById("description").value;
  let image = document.getElementById("image-upload").files;

  // untuk membuat url gambar, agar tampil
  image = URL.createObjectURL(image[0]);
  console.log(image);

  let blog = {
    title,
    content,
    image,
    duration: "1 bulan",
    author: "Malfazakki",
  };

  dataBlog.push(blog);
  console.log(dataBlog);

  renderBlog();
}

function renderBlog() {
  document.getElementById("description").innerHTML = "";

  for (let index = 0; index < dataBlog.length; index++) {
    document.getElementById("contents").innerHTML += `
        <div class="blog-list-item">
            <div class="blog-image">
                <img src="${dataBlog[index].image}" alt="blog_img" />
            </div>
            <div class="blog-content">
            <div class="btn-group">
                <button class="btn-edit">Edit Post</button>
                <button class="btn-delete">Delete Post</button>
            </div>
            <h1>
                <a href="blog-detail.html" target="_blank"
                >${dataBlog[index].title}</a
                >
            </h1>
            <div class="detail-blog-content">
                ${dataBlog[index].postAt} | ${dataBlog[index].author}
            </div>
            <p>
                ${dataBlog[index].content}
            </p>
            </div>
        </div>
    `;
  }
}
