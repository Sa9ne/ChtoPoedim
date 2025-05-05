const ingredientStages = [
  { name: 'vegan', src: '/frontend/icons/vegan.png' },
  { name: 'spicy', src: '/frontend/icons/spicy.png' },
  { name: 'hearty', src: '/frontend/icons/hearty.png' }
];

let currentStage = 0; // Текущий этап

const img = document.querySelector('.ingredient-photo');
const acceptBtn = document.querySelector('.Accept-button');
const rejectBtn = document.querySelector('.Reject-button');
const repeatBtn = document.querySelector('.Repeat-button');

function handleChoice(choice) {
  console.log(`Вы выбрали ${choice} для ${ingredientStages[currentStage].name}`);
  
  // Следующий этап
  currentStage++;
  if (currentStage < ingredientStages.length) {
      img.src = ingredientStages[currentStage].src;
  } else {
      img.src = '/frontend/icons/game-over.png'; // например, финальная картинка
      acceptBtn.style.display = 'none';
      rejectBtn.style.display = 'none';
      repeatBtn.style.display = 'inline-block';
  }
}

function resetProcess() {
  currentStage = 0;
  img.src = ingredientStages[0].src;
  acceptBtn.style.display = 'inline-block';
  rejectBtn.style.display = 'inline-block';
  repeatBtn.style.display = 'none';
}

acceptBtn.addEventListener('click', () => handleChoice('accept'));
rejectBtn.addEventListener('click', () => handleChoice('reject'));
repeatBtn.addEventListener('click', resetProcess);