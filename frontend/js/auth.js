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
// Регистрация 
async function RegisterUser(e) {
    // Не перезагружаем страницу
    e.preventDefault();

    // Вводим данные
    const name = document.getElementById("register-name").value;
    const email = document.getElementById("register-email").value;
    const password = document.getElementById("register-password").value;
    const errorEl = document.getElementById("register-error");

    try {
        // Отправляем запрос по ссылке и указываем как
        const response = await fetch("http://localhost:8081/register", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ name, email, password })
        });
            // Ждем ответа от сервера
            const data = await response.json();
        if (response.ok) {
            CloseModal();
        } else {
            errorEl.textContent = data.error || "Произошла ошибка регистрации.";
        }
    } catch (err) {
        errorEl.textContent = "Ошибка сети: " + err.message;
    }
}
// Вход
async function LogInUser(e) {
    e.preventDefault();
    
    // Вводим данные
    const email = document.getElementById("LogIn-email").value;
    const password = document.getElementById("LogIn-password").value;
    const errorEl = document.getElementById("LogIn-error");

    try {
        const response = await fetch("http://localhost:8081/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ email, password })
        });
            const data = await response.json();
        if (response.ok) {
            CloseModal();
        } else {
            errorEl.textContent = data.error || "Произошла ошибка входа."
        }
    } catch (err) {
        errorEl.textContent = "Ошибка сети: " + err.message
    }
}