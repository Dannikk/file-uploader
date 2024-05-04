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

  const url = `http://127.0.0.1:8080/download?` + new URLSearchParams({ path: fileName }).toString();

  fetch(url)
  .then(res => res.json()).then(({file, content_type}) => {

    const link = document.createElement('a');
    const div = document.createElement(('div'))
    link.href = `data:${content_type};base64,${file}`;
    link.download = fileName;
    link.textContent = `Download ${fileName}`;

    div.appendChild(link)
    link.addEventListener('click', ()=> link.remove())
    document.body.appendChild(div);
  });
}

function handleSubmit(event) {
  event.preventDefault();
  event.stopPropagation();

  uploadFiles(event.target);
}

function uploadFiles(currentForm) {
  const url = 'http://127.0.0.1:8080/upload';
  const method = 'post';

  const data = new FormData(currentForm);

  fetch(url, {
    method,    
    body: data
  })
    .then(resp => resp.json()).then(displayFileName)
}

function displayFileName(response) {
  statusMessage.textContent = response.name;
}
