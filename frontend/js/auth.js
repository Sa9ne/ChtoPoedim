// Переменная для хранения API
const API_AUTH = "http://localhost:8081";

// Функция для переключения окна входа
function toggleAuthForm() {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");

    if (authForm && registerForm) {
        if (authForm.classList.contains("show") || registerForm.classList.contains("show")) {
            authForm.classList.remove("show");
            registerForm.classList.remove("show");
        } else {
            authForm.classList.add("show");
        }
    }
}

// Функция для переключения окна регистрации
function toggleRegisterForm() {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");

    if (authForm) authForm.classList.remove("show");
    if (registerForm) registerForm.classList.add("show");
}

// Функция для переключения на форму входа
function toggleToLoginForm() {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");

    if (registerForm) registerForm.classList.remove("show");
    if (authForm) authForm.classList.add("show");
}

// Функция для регистрации пользователя
async function registerUser(event) {
    event.preventDefault();

    const username = document.querySelector("#register-form input[placeholder='Username']").value;
    const email = document.querySelector("#register-form input[placeholder='Email']").value;
    const password = document.querySelector("#register-form input[placeholder='Password']").value;

    const response = await fetch(`${API_AUTH}/register`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, email, password })
    });

    const data = await response.json();

    if (response.ok) {
        localStorage.setItem("token", data.token);
        showProfile(username);
    } else {
        alert("Ошибка регистрации: " + (data.error || "неизвестная ошибка"));
    }
}

// Функция для входа пользователя
async function loginUser(event) {
    event.preventDefault();

    const email = document.querySelector("#auth-form input[placeholder='Email']").value;
    const password = document.querySelector("#auth-form input[placeholder='Password']").value;

    const response = await fetch(`${API_AUTH}/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password })
    });

    const data = await response.json();

    if (response.ok) {
        localStorage.setItem("token", data.token);
        showProfile(data.username);
    } else {
        alert("Ошибка входа: " + (data.error || "неизвестная ошибка"));
    }
}

// Функция для отображения профиля
function showProfile(username) {
    const authForm = document.getElementById("auth-form");
    const registerForm = document.getElementById("register-form");
    const profile = document.getElementById("profile-form");
    const nickname = document.getElementById("nickname");

    if (authForm) authForm.classList.add("hidden");
    if (registerForm) registerForm.classList.add("hidden");
    if (profile) profile.classList.remove("hidden");
    if (nickname) nickname.textContent = username;
}

// Функция выхода из аккаунта
function quitAccount() {
    localStorage.removeItem("token");
    const profile = document.getElementById("profile-form");
    const authForm = document.getElementById("auth-form");

    if (profile) profile.classList.add("hidden");
    if (authForm) authForm.classList.remove("hidden");
}

// Функция для проверки авторизации
async function checkAuth() {
    const token = localStorage.getItem("token");
    if (!token) return;

    try {
        const response = await fetch(`${API_AUTH}/profile`, {
            method: "GET",
            headers: { "Authorization": `Bearer ${token}` }
        });

        const text = await response.text(); // Считываем ответ как текст
        console.log(text);  // Логируем ответ, чтобы увидеть, что приходит

        const data = JSON.parse(text); // Преобразуем в JSON, если это правильный JSON

        if (response.ok) {
            showProfile(data.username);
        } else {
            console.error("Ошибка получения данных профиля:", data);
            localStorage.removeItem("token");
        }
    } catch (error) {
        console.error("Ошибка при обработке ответа:", error);
    }
}

// Инициализация и проверка авторизации на старте
document.addEventListener("DOMContentLoaded", function () {
    checkAuth();

    const registerButton = document.getElementById("register-button");
    if (registerButton) {
        registerButton.addEventListener("click", registerUser);
    }

    const authButton = document.getElementById("auth-button");
    if (authButton) {
        authButton.addEventListener("click", loginUser);
    }

    const quitButton = document.getElementById("quit-button");
    if (quitButton) {
        quitButton.addEventListener("click", quitAccount);
    }
});
