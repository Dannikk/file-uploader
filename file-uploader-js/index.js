const form = document.querySelector('form');
form.addEventListener('submit', handleSubmit);
const fileInput = document.querySelector('input');

const statusMessage = document.getElementById('statusMessage');

const downloadButton = document.getElementById('downloadButton');
const fileNameInput = document.getElementById('fileNameInput');


downloadButton.addEventListener('click', handleDownloadClick);

function handleDownloadClick(event) {
  event.preventDefault();

  const fileName = fileNameInput.value;

  const url = `http://127.0.0.1:8080/download?` + new URLSearchParams({ path: fileName });

  fetch(url)
  .then(alert);
}


function handleSubmit(event) {
  event.preventDefault();
  event.stopPropagation();

  console.log('form submitted');

  uploadFiles(event.target);

}

function uploadFiles(currentForm) {
  const url = 'http://127.0.0.1:8080/upload';
  const method = 'post';

  const data = new FormData(currentForm);debugger

  fetch(url, {
    method,    
    body: data
  })
    .then(resp => resp.json()).then(displayFileName)
}

function displayFileName(response) {
  // const display = document.createElement('displayName');

  // // display.href = response.name;
  // display.textContent = response.name;

  // document.body.appendChild(display);

  statusMessage.textContent = response.name;
}

function updateStatusMessage(text) {
  statusMessage.textContent = text;
}

function saveFile(file) {
  const a = document.createElement('a');
  a.href = file.files;

  // a.download = "new_file.txt";
  // a.textContent = "Download";
  a.textContent = file.files.file

  document.body.appendChild(a);
}
