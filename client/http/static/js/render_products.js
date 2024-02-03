function renderProducts(products, pageEl) {

  products.forEach((p) => {
    const html = ` 
        <li>
        
        <h2>${p.title}</h2>
        <p><strong>Site:</strong> ${p.site}</p>
        <p><strong>Price:</strong> ${p.price}</p>
        <p><strong>URL:</strong> <a href="${p.url}" target="_blank">${p.url}</a></p>
        <img src=${p.image} alt=${p.title}>
         
        </li>
                 `;

    pageEl.innerHTML += html;
  });
}
