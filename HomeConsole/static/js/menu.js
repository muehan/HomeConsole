$(function(){
    'use strict'

    var currentHref = window.location.pathname.toLowerCase().replace("/", "");

    if(!currentHref){
        currentHref = "overview";
    }
    
    var ulElements = document.getElementsByClassName("nav");

    var oldMenuActive = document.getElementsByClassName("active");
    oldMenuActive.className = "";
    
    for (var x = 0; x < ulElements.length; x++) {
        for (var i = 0; i < ulElements[x].childNodes.length; i++) {
            var liElement = ulElements[x].childNodes[i];
            if(liElement.innerText){
                if(liElement.innerText.trim().toLowerCase().replace(" ", "") == currentHref){
                    liElement.className += " active";
                    break;
                }
            }
        }
    }
});