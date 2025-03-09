document.addEventListener("DOMContentLoaded", function () {
    const searchForm = document.querySelector(".search-container form");
    const searchInput = document.querySelector(".search-input");
    const resultsContainer = document.querySelector(".main-cont");

    searchForm.addEventListener("submit", function (event) {
        event.preventDefault(); // Отменяем стандартную отправку формы

        const query = searchInput.value.trim();
        if (!query) {
            resultsContainer.innerHTML = "<p>Введите поисковый запрос</p>";
            return;
        }

    fetch(`/search?search=${encodeURIComponent(query)}`)
        .then(response => response.json())
        .then(data => {
            console.log("Ответ от сервера:", data); // Проверяем JSON в консоли

            if (data.message) {
                resultsContainer.innerHTML = `<p>${data.message}</p>`;
            } else {
                let dishesList = "<ul>";
                data.forEach(dish => {
                    // Проверяем, есть ли time в объекте
                    let timeInfo = dish.time ? `${dish.time} минут` : "время не указано";

                    dishesList += `<li>
                        <strong>${dish.name}</strong><br>
                        ${dish.description}<br>
                        Время приготовления: ${timeInfo}
                    </li>`;
                });
                dishesList += "</ul>";
                resultsContainer.innerHTML = dishesList;
            }
        })
        .catch(error => {
            console.error("Ошибка:", error);
            resultsContainer.innerHTML = "<p>Ошибка при загрузке данных</p>";
        });
    });
});
