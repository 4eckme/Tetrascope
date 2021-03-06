// Tetrascope fractal animation
// Availiable to execute in browser console
// For base tetrascope image set R=1024, n=1
(function(){ var R = 256; var D=2*R+1;
var pow = Math.pow; var floor=Math.floor; var ceil=Math.ceil; var log2 = Math.log2;
var rgb = function(x, y, n){ return floor(   (x*x*n+y*y*n)*pow(16, ( 6-ceil(log2(x*x*n+y*y*n)/4)) )   ); }
document.body.innerHTML=('<canvas id="C" width="'+D+'" height="'+D+'"></canvas>');
var canvas = document.getElementById('C');
var ctx = canvas.getContext('2d');
window.n=0; var si = setInterval(function () {
    window.n += 1;
    for(var x = 0;x<=R;x++) {
      for(var y = 0;y<=R;y++) {
        ctx.fillStyle='#'+rgb(x, y, n).toString(16);                    
        if (x===0&&y===0) ctx.fillStyle="#000000";
        var X1 = R-x;
        var Y1 = R-y;
        var X2 = R+x;
        var Y2 = R+y;
        if (  ( x*x+y*y ) < R*R ) {      
          ctx.fillRect(X1, Y1, 1, 1);
          ctx.fillRect(X1, Y2, 1, 1);
          ctx.fillRect(X2, Y1, 1, 1);
          ctx.fillRect(X2, Y2, 1, 1);
}   }   }   }, 100);
})();
