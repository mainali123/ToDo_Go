const passwords= document.querySelectorAll('#password');
const showPw= document.querySelectorAll('#showpass');
const email= document.getElementById('email');


showPw.forEach((show, index)=>{

    show.addEventListener('click', ()=>{
        if(passwords[index].type==='password'){
            passwords[index].type='text';
            show.src='../static/images/eye-slash-solid.svg'
        } else{
            passwords[index].type='password';
            show.src='../static/images/viuew.png';
        }
    })
});



email.addEventListener('blur', () => { 
    const emailValue = email.value;
    const emailRegex = /^\S+@\S+\.\S+$/;

    if (emailRegex.test(emailValue)) {
        email.style.border = '1px solid green';
    } else {
        email.style.border = '1px solid red';
    }
});