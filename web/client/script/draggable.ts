let dash = <HTMLDivElement>document.getElementById("dashboard");
let new_x: number;
let new_y: number;

dash.addEventListener("drag", (e: MouseEvent) => {
    let target = <HTMLElement>e.target;
    
    let tx = target.offsetLeft;
    let ty = target.offsetTop;
    let mx = tx + e.offsetX;
    let my = ty + e.offsetY;
    
    if (
        mx > 0 &&
        mx < dash.offsetWidth &&
        my > 0 &&
        my < dash.offsetHeight
    ) {
        new_x = mx;
        new_y = my;
    }
});

dash.addEventListener("mouseup", (e: MouseEvent) => {
    let target = <HTMLElement>e.target;
});
