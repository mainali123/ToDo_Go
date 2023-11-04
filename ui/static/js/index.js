const signIn = document.getElementById('signin');
const overlay = document.querySelector('.overlay');
const popUp = document.querySelector('.popUp');
const exit = document.querySelector('.img');

signIn.addEventListener('click', (e)=>{
    e.preventDefault();
    overlay.classList.add('click');
    popUp.classList.add('active');
})
