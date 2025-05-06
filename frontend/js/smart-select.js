const ingredientStages = [
  { name: 'vegan', src: '/frontend/icons/vegan.png' },
  { name: 'spicy', src: '/frontend/icons/spicy.png' },
  { name: 'hearty', src: '/frontend/icons/hearty.png' }
];

let currentStage = 0;
let userChoices = [];

const img = document.querySelector('.ingredient-photo');
const acceptBtn = document.querySelector('.Accept-button');
const rejectBtn = document.querySelector('.Reject-button');
const repeatBtn = document.querySelector('.Repeat-button');

// Вызывается при выборе accept/reject
function handleChoice(choice) {
  const isAccepted = (choice === 'accept');

  // Сохраняем выбор
  userChoices.push({
    stage: ingredientStages[currentStage].name,
    choice: isAccepted
  });

  currentStage++;

  if (currentStage < ingredientStages.length) {
    img.src = ingredientStages[currentStage].src;
  } else {
    // Все этапы пройдены
    img.src = '/frontend/icons/game-over.png';
    acceptBtn.style.display = 'none';
    rejectBtn.style.display = 'none';
    repeatBtn.style.display = 'inline-block';

    fetch('http://localhost:8082/SmartSelectFood', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(userChoices)
    })
    .then(response => {
      if (!response.ok) throw new Error('Ошибка при отправке');
      return response.json(); // читаем ответ как JSON
    })
    .then(dishes => {
      const container = document.getElementById('smart-card-place');
      const template = document.getElementById('Smart-catalog');
    
      // Очищаем старые карточки, если они есть
      container.innerHTML = '';
    
      // Проходим по массиву блюд
      dishes.forEach(dish => {
        const card = template.content.cloneNode(true);

        // Вставляем данные
        card.querySelector('.name-for-smart-card').textContent = dish.name;

        const img = card.querySelector('.img-smart-select');
        img.alt = dish.name;
        // Заменим иконку, если есть картинка блюда:
        if (dish.image) {
          img.src = dish.image;
        } else {
          img.src = ''; 
        }
    
        container.appendChild(card);
      });
    })
    .catch(err => console.error('Сетевая ошибка:', err));    
  }
}

// Сброс всех данных и перезапуск
function resetProcess() {
  currentStage = 0;
  userChoices = [];
  img.src = ingredientStages[0].src;
  acceptBtn.style.display = 'inline-block';
  rejectBtn.style.display = 'inline-block';
  repeatBtn.style.display = 'none';

  // Удаляем старые карточки
  const container = document.getElementById('smart-card-place');
  container.innerHTML = '';
}

// Обработчики событий
acceptBtn.addEventListener('click', () => handleChoice('accept'));
rejectBtn.addEventListener('click', () => handleChoice('reject'));
repeatBtn.addEventListener('click', resetProcess);