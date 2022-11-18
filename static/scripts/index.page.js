~~function(){
    var content = document.querySelector(".gox-body-content .container-fluid")
    var fragment = document.createDocumentFragment()
    for(let i=0;i<100;i++){
        let h1Title = document.createElement("h1")
        h1Title.innerText = "Goxlog Hello!"
        fragment.appendChild(h1Title)
    }
    content.appendChild(fragment)
    console.log("success.")
}()