// To open up the burger menu
// The following code is based off a toggle menu by @Bradcomp
// source: https://gist.github.com/Bradcomp/a9ef2ef322a8e8017443b626208999c1
(function() {
    var burger = document.querySelector('.burger');
    var menu = document.querySelector('#'+burger.dataset.target);
    burger.addEventListener('click', function() {
        burger.classList.toggle('is-active');
        menu.classList.toggle('is-active');
    });
})();

//This supposedly solves the form resubmission issue with the lang button.
//From https://dtbaker.net/files/prevent-post-resubmit.php
if (window.history.replaceState) {
    window.history.replaceState(null, null, window.location.href);
}