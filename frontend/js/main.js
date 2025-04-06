// Константы
    window.AuthForm = document.getElementById("auth-form")
    window.RegisterForm = document.getElementById("register-form")
    window.OpenLang = document.getElementById("lang-modal")
    window.OpenProf = document.getElementById("profile-form")
// Функция закрытия всех модальных окон
function CloseModal(expect = null) {
    const modals = [AuthForm, RegisterForm, OpenLang];

    modals.forEach(modal => {
        if (modal !== expect) {
            modal.classList.add("hidden");
        }
    });
}