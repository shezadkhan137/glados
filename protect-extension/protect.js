function cookieinfo(cookies, callback){

    console.log("Cookie info called", cookies)
    var inserted = 0;

    for (var i=0; i < cookies.length; i++){
         d = cookies[i]
         data = {"name": d.Name, "url": "https://" + d.Domain + d.Path, "value": d.Value}
         console.log("Setting cookie data", data)
         chrome.cookies.set(data,function (cookie){
            console.log(JSON.stringify(cookie));
            console.log(chrome.extension.lastError);
            console.log(chrome.runtime.lastError);

         if (++inserted == cookies.length) {
             console.log("All cookies set")
            callback()
          }
        });
    }

}

chrome.webNavigation.onCompleted.addListener(function(details){
    console.log("Got webNavigation onCompleted")
    chrome.tabs.get(details.tabId, function(tab){
        console.log(tab)
        console.log("Got tab", tab)
        var parser = document.createElement('a');
        parser.href = tab.url

        if (parser.host === "lmet.aiwip.com"){
            var xhr = new XMLHttpRequest();
            xhr.open("GET", "http://localhost:8080/cookies", true);
            xhr.onreadystatechange = function(){
                if (xhr.readyState == 4) {
                    result = JSON.parse(xhr.responseText)
                    console.log(result)
                    cookieinfo(result, function(){
                        var code = 'window.location.reload();'; 
                        chrome.tabs.executeScript(tab.id, {code: code});
                    })
                }
            }
            xhr.send();
        }
    })
})
