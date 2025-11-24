// Mobile menu toggle
const mobileMenuButton = document.getElementById('mobile-menu-button');
const mobileMenu = document.getElementById('mobile-menu');

if (mobileMenuButton && mobileMenu) {
    mobileMenuButton.addEventListener('click', () => {
        mobileMenu.classList.toggle('hidden');
    });

    // Close mobile menu when clicking on a link
    const mobileMenuLinks = mobileMenu.querySelectorAll('a');
    mobileMenuLinks.forEach((link) => {
        link.addEventListener('click', () => {
            mobileMenu.classList.add('hidden');
        });
    });

    // Close mobile menu when clicking outside
    document.addEventListener('click', (e) => {
        if (!mobileMenu.contains(e.target) && !mobileMenuButton.contains(e.target)) {
            mobileMenu.classList.add('hidden');
        }
    });
}

const contactForm = document.querySelector('.contact-form');
if (contactForm) {
    contactForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const submitButton = contactForm.querySelector('.btn-submit');
        const originalText = submitButton.textContent;
        
        submitButton.textContent = 'Отправка...';
        submitButton.disabled = true;

        try {
            const formData = new URLSearchParams(); // эта функция возвращает обычную форму, а не multiPart, что собственно и нужно
            formData.append('name', contactForm.querySelector('[name="name"]').value);
            formData.append('email', contactForm.querySelector('[name="email"]').value);
            formData.append('phone', contactForm.querySelector('[name="phone"]').value);
            formData.append('agreement', contactForm.querySelector('[name="agreement"]').checked ? 'on' : '');
            
            const response = await fetch('/contact', {
                method: 'POST',
                body: formData
            });

            if (response.ok) {
                alert('Спасибо за ваше сообщение! Мы свяжемся с вами в ближайшее время.');
                contactForm.reset();
            } else {
                throw new Error('Ошибка сервера');
            }
        } catch (error) {
            console.error('Ошибка:', error);
            alert('Произошла ошибка при отправке формы. Пожалуйста, попробуйте еще раз.');
        } finally {
            // Восстанавливаем кнопку
            submitButton.textContent = originalText;
            submitButton.disabled = false;
        }
    });
}

// Smooth scroll for anchor links
document.querySelectorAll('a[href^="#"]').forEach(anchor => {
    anchor.addEventListener('click', function (e) {
        e.preventDefault();
        const target = document.querySelector(this.getAttribute('href'));
        if (target) {
            const offset = 70;
            const targetPosition = target.offsetTop - offset;
            window.scrollTo({
                top: targetPosition,
                behavior: 'smooth'
            });
        }
    });
});