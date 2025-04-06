// Открытие окна входа
function toggleLogInForm() {
    CloseModal(AuthForm);
    AuthForm.classList.toggle("hidden");
}
// Открытие регистрационного окна
function toggleRegisterForm() {
    CloseModal(RegisterForm);
    RegisterForm.classList.toggle("hidden"); 
}
