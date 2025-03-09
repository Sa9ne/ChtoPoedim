// Функция для переключения окна входа
function toggleAuthForm() {
    var authForm = document.getElementById("auth-form");
    authForm.classList.toggle("show"); // Показываем/скрываем форму входа
}

// Функция для переключения окна регистрации
function toggleRegisterForm() {
    var registerForm = document.getElementById("register-form");
    registerForm.classList.toggle("show"); // Показываем/скрываем форму регистрации
}

// Переход из окна входа в окно регистрации
function toggleToRegisterForm() {
    var authForm = document.getElementById("auth-form");
    var registerForm = document.getElementById("register-form");
    authForm.classList.remove("show"); // Скрываем окно входа
    registerForm.classList.add("show"); // Показываем окно регистрации
}

// Переход из окна регистрации в окно входа
function toggleToLoginForm() {
    var registerForm = document.getElementById("register-form");
    var authForm = document.getElementById("auth-form");
    registerForm.classList.remove("show"); // Скрываем окно регистрации
    authForm.classList.add("show"); // Показываем окно входа
}

// Открытие модального окна для входа (например, при нажатии на кнопку)
document.querySelector("#open-auth-form").addEventListener("click", toggleAuthForm);
