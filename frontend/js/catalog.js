// Функция открытия каталога
function OpenCatalog() {
  CloseModal(Catalog);
  Catalog.classList.toggle("hidden")
}
// Функция открытия окна блюда дня
function OpenDayDish() {
  CloseModal(DishDay);
  DishDay.classList.toggle("hidden")
  DayDish();
}
// Функция открытия окна выбора блюд
function OpenSmartSelect() {
  CloseModal();
  if (!SmartSelect.classList.contains("hidden")) {
    CloseMainWind(InfoSite);
    InfoSite.classList.remove("hidden");
  } else {
    CloseMainWind(SmartSelect);
    SmartSelect.classList.remove("hidden");
  }
}
// Функция открытия каталога еды
function OpenFoodCatalog() {
  CloseModal();
  if (!FoodCatalog.classList.contains("hidden")) {
    CloseMainWind(InfoSite);
    InfoSite.classList.remove("hidden");
  } else {
    CloseMainWind(FoodCatalog);
    FoodCatalog.classList.remove("hidden");
  }
}
// Кнопка возвращения на основной экран
function BackToMain() {
  if (!SmartSelect.classList.contains("hidden")) {
    SmartSelect.classList.add("hidden")
  }

  if (!FoodCatalog.classList.contains("hidden")) {
    FoodCatalog.classList.add("hidden")
  }

  InfoSite.classList.toggle("hidden")
}
// Функция загрузки перевода на русский
let translations_dish = {};

async function LoadTranslation(lang) {
  try {
    const response = await fetch(`/frontend/lang/dish/${lang}.json`);
    if (!response.ok) throw new Error("Ошибка загрузки перевода");
    translations_dish = await response.json();
  } catch (error) {
    console.error("Loading translation problem")
  }
}
// Функция перевода на русский
function translate(text) {
  return translations_dish[text] || text;
}
// Функция блюда дня
async function DayDish() {
  try {
    const response = await fetch("http://localhost:8082/DishDay");
    if (!response.ok) throw new Error("Error found dish")

    const lang = localStorage.getItem("language")
    await LoadTranslation(lang);

		const dish = await response.json();

		// Заполняем содержимое
		DishName.textContent = translate(dish.name) || "No name";
		DishDescription.textContent = translate(dish.description) || "Null description";
		DishTime.textContent = dish.time || "-";
  }  catch (error) {
		console.error("Ошибка:", error);
	}
}
// Функция вывода каталога
document.addEventListener("DOMContentLoaded", () => {
  fetch("http://localhost:8082/DishCatalog")
    .then(response => response.json())
    .then(data => {
      const container = document.getElementById("cards-dish");
      const template = document.getElementById("Dish-template");

      if (!Array.isArray(data) || data.length === 0) {
        console.warn("Нет блюд для отображения");
        return;
      }

      data.forEach(dish => {
        const clone = template.content.cloneNode(true);

        const img = clone.querySelector(".dish-ctl-img");
        const name = clone.querySelector(".dish-ctl-name");
        const favBtn = clone.querySelector(".dish-fav-btn");

        img.alt = dish.name;
        name.textContent = dish.name;
        favBtn.setAttribute("data-id", dish.id); // для обработки "избранного"

        container.appendChild(clone);
      });
    })
    .catch(error => {
      console.error("Ошибка при получении блюд:", error);
    });
});
