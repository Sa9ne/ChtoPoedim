// Константы
    window.authForm = document.getElementById("auth-form")
    window.RegisterForm = document.getElementById("register-form")
    window.OpenLang = document.getElementById("lang-modal")
// Функция закрытия всех модальных окон
function CloseModal() {
    authForm.classList.add("hidden");
    RegisterForm.classList.add("hidden");
    OpenLang.classList.add("hidden");
}