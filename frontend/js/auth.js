// Константы
    const authForm = document.getElementById("auth-form")
    const RegisterForm = document.getElementById("register-form")
// Открытие окна входа
function toggleLogInForm() {
    RegisterForm.style.display = "none";

    if (authForm.style.display === "none" || authForm.style.display === "") {
        authForm.style.display = "block"
    } else {
        authForm.style.display = "none"
    }
}
// Открытие регистрационного окна
function toggleRegisterForm() {
    authForm.style.display = "none";

    if (RegisterForm.style.display === "none" || RegisterForm.style.display === "") {
        RegisterForm.style.display = "block"
    } else {
        RegisterForm.style.display = "none"
    }
}
