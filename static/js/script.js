const url = "ws://" + window.location.host + "/sock";

const messengerArea = document.getElementById('chatBox')
const ws = new WebSocket(url)
try {
    ws.onopen = () => {
        console.log("соединение открыто")
    }
    ws.onmessage = (m) => {
        let div = document.createElement('div')
        div.classList.add('messenger__message')
        div.textContent = m.data
        messengerArea.appendChild(div)
    }
    ws.onerror = (e) => {
        console.log("Ошибка отправки :" + e.data)
    }
    ws.onclose = (m) => {
        console.log("Соединение разорвано :" + m.data)
    }
} catch (e) {
    console.log(e)
}

const inputName = document.getElementById('inputName')
const saveNameBtn = document.getElementById('saveNameBtn')

let isSaveName = false
saveNameBtn.addEventListener('click', event => {

    if (inputName.value === '') {
        return
    }

    isSaveName = true
    inputName.setAttribute('disabled', 'disabled')
    saveNameBtn.style.display = 'none'
    errorsBlock.style.display = 'none'
})

const inputMsg = document.getElementById('inputMsg')
const btnSend = document.getElementById('send')

const errorsBlock = document.querySelector('.messenger__error')

    btnSend.addEventListener('click', event => {

        if (inputMsg.value === '') {
            return
        }

        if (isSaveName) {
            ws.send(inputName.value + ' : ' + inputMsg.value)
            inputMsg.value = ""
        } else {
            errorsBlock.style.color = 'red'
            errorsBlock.textContent = 'Введите имя!'
        }

    })