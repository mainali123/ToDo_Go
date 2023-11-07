//password show hide script

const showPassword = document.querySelectorAll(".password-show");

showPassword.forEach((show) => {
	show.addEventListener("click", () => {
		const password = show.parentElement.children[1];
		if (password.type === "password") {
			password.type = "text";
			show.innerHTML = '<i class="fa-solid fa-eye-slash"></i>';
		} else {
			password.type = "password";
			show.innerHTML = '<i class="fa-solid fa-eye"></i>';
		}
	});
});

// alert message script

const alert = document.querySelector(".alert-box-dialog");
const alertMessage = document.querySelector(".alert-message-dialog");

function showAlert(message) {
	alertMessage.innerHTML = message;
	alert.style.visibility = "visible";
	alert.style.opacity = "1";

	setTimeout(() => {
		alert.style.visibility = "hidden";
		alert.style.opacity = "0";
	}, 5000);
}

//form check script

const email = document.querySelector(".login-email");
const password = document.querySelector(".login-password");
const loginForm = document.querySelector(".login-form");

loginForm.addEventListener("submit", (e) => {
	e.preventDefault();

	if (email.value === "") {
		email.focus();
		showAlert("Please fill all the fields");
	} else if (password.value === "") {
		password.focus();
		showAlert("Please fill all the fields");
	} else {
		showAlert("The login system is not ready yet");
	}
});
