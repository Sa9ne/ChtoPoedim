// Константы
    const authForm = document.getElementById("auth-form")
    const RegisterForm = document.getElementById("register-form")
// Открытие окна входа
function toggleLogInForm() {
    RegisterForm.classList.add("hidden"); 

    authForm.classList.toggle("hidden"); 
}
// Открытие регистрационного окна
function toggleRegisterForm() {
    authForm.classList.add("hidden"); 

    RegisterForm.classList.toggle("hidden"); 
}
