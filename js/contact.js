function submitData() {
  let name = document.getElementById("name").value;
  let email = document.getElementById("email").value;
  let phoneNumber = document.getElementById("phone-number").value;
  let subject = document.getElementById("subject").value;
  let messages = document.getElementById("messages").value;

  if (
    name === "" ||
    email === "" ||
    phoneNumber === "" ||
    subject === "" ||
    messages === ""
  ) {
    return alert("Pastikan semua kolom formulir terisi");
  }

  let emailReceiver = "malfazakki@gmail.com";

  let aEmail = document.createElement("a");
  aEmail.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, nama saya ${name}, ${messages}. Silakan kontak saya di nomor ${phoneNumber}, terima kasih.`;
  aEmail.click();

  console.log(name);
  console.log(email);
  console.log(phoneNumber);
  console.log(subject);
  console.log(messages);

  let emailer = {
    name,
    email,
    phoneNumber,
    subject,
    messages,
  };

  console.log(emailer);
}
