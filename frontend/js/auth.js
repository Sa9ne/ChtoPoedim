// Функция для открытия и закрытия окна входа
function toggleAuthForm() {
    const authForm = document.getElementById("auth-form");
    authForm.classList.toggle("hidden");
}

// Функция для переключения на форму регистрации
function toggleToRegisterForm() {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");

    // Скрываем форму входа и показываем форму регистрации
    if (authForm) authForm.classList.remove("show");
    if (registerForm) registerForm.classList.add("show");
}

// Функция для переключения на форму входа
function toggleToLoginForm() {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");

    // Скрываем форму регистрации и показываем форму входа
    if (registerForm) registerForm.classList.remove("show");
    if (authForm) authForm.classList.add("show");
}

// Функция для закрытия окна регистрации
function toggleRegisterForm() {
    const registerForm = document.getElementById("register-form");

    // Закрыть окно регистрации
    if (registerForm) registerForm.classList.remove("show");
}

// Функция для регистрации нового пользователя
async function registerUser(event) {
    event.preventDefault();

    const username = document.querySelector("#register-form input[placeholder='Username']").value;
    const email = document.querySelector("#register-form input[placeholder='Email']").value;
    const password = document.querySelector("#register-form input[placeholder='Password']").value;

    try {
        const response = await fetch("http://localhost:8081/register", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ username, email, password })
        });

        const data = await response.json();

        if (response.ok) {
            alert("Регистрация прошла успешно!");
            toggleToLoginForm(); // Переключаем на форму входа
        } else {
            alert("Ошибка регистрации: " + data.error);
        }
    } catch (error) {
        console.error("Ошибка регистрации:", error);
    }
}

// Функция для входа пользователя
async function loginUser(event) {
    event.preventDefault();

    const email = document.querySelector("#auth-form input[placeholder='Email']").value;
    const password = document.querySelector("#auth-form input[placeholder='Password']").value;

    try {
        const response = await fetch("http://localhost:8081/login", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ email, password })
        });

        const data = await response.json();

        if (response.ok) {
            localStorage.setItem("auth_token", data.token);
            showProfile(data.username); // Если авторизация успешна, показываем профиль
        } else {
            alert("Ошибка входа: " + data.error);
        }
    } catch (error) {
        console.error("Ошибка входа:", error);
    }
}

// Функция для отображения профиля
function showProfile(username) {
    const profileForm = document.getElementById("profile-form");
    const authForm = document.getElementById("auth-form");

    if (authForm) authForm.classList.add("hidden"); // Прячем форму входа, если пользователь авторизован
    if (profileForm) profileForm.classList.remove("hidden"); // Показываем окно профиля

    const nickname = document.getElementById("nickname");
    if (nickname) {
        nickname.textContent = username; // Отображаем никнейм пользователя в окне профиля
    }
}

// Функция для выхода пользователя
function quitAccount() {
    localStorage.removeItem("auth_token"); // Удаляем токен из localStorage
    toggleProfileForm(); // Закрываем окно профиля
    toggleAuthForm(); // Открываем окно входа
}

// Функция для открытия и закрытия окна профиля
function toggleProfileForm() {
    const profileForm = document.getElementById("profile-form");
    profileForm.classList.toggle("hidden");
}

// Инициализация слушателей событий для форм
document.addEventListener("DOMContentLoaded", function() {
    const token = localStorage.getItem("auth_token");

    if (token) {
        fetch("http://localhost:8081/profile", {
            method: "GET",
            headers: {
                "Authorization": `Bearer ${token}`
            }
        })
        .then(response => response.json())
        .then(data => {
            if (data.username) {
                showProfile(data.username); // Если токен есть и данные профиля получены, показываем профиль
            }
        })
        .catch(error => console.error("Ошибка при получении профиля:", error));
    }

    const registerButton = document.getElementById("register-button");
    const authButton = document.getElementById("auth-button");

    if (registerButton) {
        registerButton.addEventListener("click", registerUser);
    }

    if (authButton) {
        authButton.addEventListener("click", loginUser);
    }
});