function kyakushitsu() {
    let table = document.getElementById('syussohyo');

    for (let row of table.rows) {
        for(let cell of row.cells) {
            if (cell.innerText === '1') {
                if (cell.className === 'non-leg-type') {
                    cell.className = 'string-center leg-type';
                    console.log(cell.innerText);
                }
            }
        }
    }
}

function ailank() {
    let table = document.getElementById('syussohyo');

    for (let row of table.rows) {
        for(let cell of row.cells) {
            if (cell.innerText === '1') {
                if (cell.className === 'myai-ranking') {
                    cell.className = 'string-center myai-ranking top-ai-lank';
                    console.log(cell.innerText);
                }
            }
            if (cell.innerText === '2') {
                if (cell.className === 'myai-ranking') {
                    cell.className = 'string-center myai-ranking second-ai-lank';
                    console.log(cell.innerText);
                }
            }
            if (cell.innerText === '3') {
                if (cell.className === 'myai-ranking') {
                    cell.className = 'string-center myai-ranking thrid-ai-lank';
                    console.log(cell.innerText);
                }
            }
        }
    }
}

function mywaku() {
    let table = document.getElementById('syussohyo');
    var count = document.getElementById('horse-count');
    count = Number(count.textContent)

    console.log('頭数：', count)
    let i = 0;
    for (let row of table.rows) {
        if ( i < count *0.30) {
            for(let cell of row.cells) {
                cell.className = 'string-center mywaku-bg-inner'
                break; // 1回目で抜ける
            }
        } else if ( i < count * 0.75) {
            for(let cell of row.cells) {
                cell.className = 'string-center mywaku-bg-middle'
                break; // 1回目で抜ける
            }
        } else {
            for(let cell of row.cells) {
                cell.className = 'string-center mywaku-bg-outer'
                break; // 1回目で抜ける
            }
        }
        i++;
    }
}

function mainRace() {

    let table = document.getElementById('todays-race');
    let regex1 = /G1/g;
    let regex2 = /G2/g;
    let regex3 = /G3/g;

    for (let row of table.rows) {
        for(let cell of row.cells) {
            let result = cell.innerText.search(regex1);
            if (result !== -1) {
                cell.className = 'g1-race';
                console.log(cell.innerText);
                break;
            }
            result = cell.innerText.search(regex2);
            if (result !== -1) {
                cell.className = 'g2-race';
                console.log(cell.innerText);
                break;
            }
            result = cell.innerText.search(regex3);
            if (result !== -1) {
                cell.className = 'g3-race';
                console.log(cell.innerText);
                break;
            }
       }
    }
}