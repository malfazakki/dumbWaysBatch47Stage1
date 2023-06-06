function submitData() {
  let name = document.getElementById("input-name").value;
  let email = document.getElementById("input-email").value;
  let phone = document.getElementById("input-phone").value;
  let subject = document.getElementById("input-subject").value;
  let message = document.getElementById("input-message").value;

  if (name == "" || email == "" || phone == "" || subject == "" || message == "") {
    return alert("Pastikan Semua Kolom Terisi!");
  }

  let emailReceiver = "malfazakki@gmail.com";

  let emailLink = document.createElement("a");
  emailLink.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name}, ${message}. Mohon hubungi saya di nomor ini: ${phone}, Terima Kasih.`;
  emailLink.click();

  let emailer = {
    name,
    email,
    phone,
    subject,
    message,
  };

  console.log(emailer);
  resetForm();
}

function resetForm() {
  document.getElementById("myForm").reset();
}
