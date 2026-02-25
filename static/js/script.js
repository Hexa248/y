function openPopup(title, img, desc, id) {
  const popup = document.getElementById('popup');
  if (!popup) return;
  document.getElementById('ptitle').innerText = title;
  document.getElementById('pimg').src = img;
  document.getElementById('pdesc').innerText = desc;
  document.getElementById('pdetail').href = '/hotel?id=' + id;
  popup.classList.remove('hidden');
}

function closePopup() {
  const popup = document.getElementById('popup');
  if (!popup) return;
  popup.classList.add('hidden');
}

function filterCards(selector, attrName, keyword) {
  const cards = document.querySelectorAll(selector);
  let visible = 0;
  cards.forEach((card) => {
    const name = (card.getAttribute(attrName) || '').toLowerCase();
    const show = !keyword || name.includes(keyword.toLowerCase());
    card.style.display = show ? '' : 'none';
    if (show) visible += 1;
  });
  return visible;
}

function closeAllSearchPopups(shell) {
  shell.querySelectorAll('.search-popup').forEach((p) => p.classList.add('hidden'));
}

function initSearchShell(shell) {
  const mode = shell.dataset.searchMode;
  const cityUrl = shell.dataset.cityUrl || '/city?name=';

  const destinationInput = shell.querySelector('.search-destination');
  const destinationPopup = shell.querySelector('.destination-popup');
  const optionButtons = shell.querySelectorAll('.search-option');
  const searchSubmit = shell.querySelector('.search-submit');

  const dateTrigger = shell.querySelector('.search-date-trigger');
  const guestTrigger = shell.querySelector('.search-guest-trigger');
  const datePopup = shell.querySelector('.date-popup');
  const guestPopup = shell.querySelector('.guest-popup');

  const checkinInput = shell.querySelector('.checkin-input');
  const checkoutInput = shell.querySelector('.checkout-input');
  const adultInput = shell.querySelector('.adult-input');
  const roomInput = shell.querySelector('.room-input');
  const dateSummary = shell.querySelector('.date-summary');
  const guestSummary = shell.querySelector('.guest-summary');

  function applyDestinationFilter() {
    const keyword = destinationInput.value.trim().toLowerCase();

    optionButtons.forEach((btn) => {
      const match = btn.innerText.toLowerCase().includes(keyword);
      btn.style.display = match ? '' : 'none';
    });

    if (mode === 'city') {
      filterCards('.city-tile', 'data-city-name', keyword);
    }
    if (mode === 'hotel') {
      const visible = filterCards('.hotel-card', 'data-hotel-name', keyword);
      const empty = document.getElementById('emptyHotelResult');
      if (empty) empty.classList.toggle('hidden', visible > 0);
    }
  }

  destinationInput.addEventListener('focus', () => {
    closeAllSearchPopups(shell);
    destinationPopup.classList.remove('hidden');
    applyDestinationFilter();
  });

  destinationInput.addEventListener('input', () => {
    destinationPopup.classList.remove('hidden');
    applyDestinationFilter();
  });

  destinationInput.addEventListener('keydown', (e) => {
    if (e.key === 'Enter') {
      e.preventDefault();
      searchSubmit.click();
    }
  });

  optionButtons.forEach((btn) => {
    btn.addEventListener('click', () => {
      destinationInput.value = btn.innerText.trim();
      destinationPopup.classList.add('hidden');
      applyDestinationFilter();
      if (mode === 'city') {
        window.location.href = cityUrl + encodeURIComponent(destinationInput.value.trim());
      }
    });
  });

  dateTrigger?.addEventListener('click', () => {
    closeAllSearchPopups(shell);
    datePopup.classList.remove('hidden');
  });

  guestTrigger?.addEventListener('click', () => {
    closeAllSearchPopups(shell);
    guestPopup.classList.remove('hidden');
  });

  shell.querySelector('.apply-date')?.addEventListener('click', () => {
    const checkin = checkinInput?.value;
    const checkout = checkoutInput?.value;
    if (checkin && checkout) {
      dateSummary.textContent = `${checkin} - ${checkout}`;
    }
    datePopup.classList.add('hidden');
  });

  shell.querySelector('.apply-guest')?.addEventListener('click', () => {
    const adults = adultInput?.value || 2;
    const rooms = roomInput?.value || 1;
    guestSummary.textContent = `${adults} Dewasa, ${rooms} Kamar`;
    guestPopup.classList.add('hidden');
  });

  searchSubmit?.addEventListener('click', () => {
    const keyword = destinationInput.value.trim();
    if (!keyword) return;

    if (mode === 'city') {
      const exact = Array.from(optionButtons).find((btn) => btn.innerText.trim().toLowerCase() === keyword.toLowerCase());
      const candidate = exact || Array.from(optionButtons).find((btn) => btn.style.display !== 'none');
      if (candidate) {
        window.location.href = cityUrl + encodeURIComponent(candidate.innerText.trim());
      }
      return;
    }

    applyDestinationFilter();
    const firstVisible = document.querySelector('.hotel-card:not([style*="display: none"])');
    if (firstVisible) {
      firstVisible.scrollIntoView({ behavior: 'smooth', block: 'center' });
    }
  });

  document.addEventListener('click', (e) => {
    if (!shell.contains(e.target)) {
      closeAllSearchPopups(shell);
    }
  });
}

document.querySelectorAll('.search-shell').forEach(initSearchShell);

document.addEventListener('click', (e) => {
  const popup = document.getElementById('popup');
  if (popup && !popup.classList.contains('hidden') && e.target.id === 'popup') closePopup();
});

document.addEventListener('keydown', (e) => {
  if (e.key === 'Escape') {
    closePopup();
    document.querySelectorAll('.search-shell').forEach((shell) => closeAllSearchPopups(shell));
  }
});


function initDockedSearchAndNavbar() {
  const topbar = document.querySelector('.topbar');
  const shell = document.querySelector('.search-shell');
  if (!topbar && !shell) return;

  let lastY = window.scrollY;
  let shellTop = 0;
  let shellHeight = 0;
  let spacer = null;

  if (shell) {
    spacer = document.createElement('div');
    spacer.className = 'search-shell-placeholder';
    shell.parentNode.insertBefore(spacer, shell);
  }

  function recalc() {
    if (!shell) return;
    shell.classList.remove('is-docked');
    spacer.classList.remove('active');
    spacer.style.height = '0px';
    const rect = shell.getBoundingClientRect();
    shellTop = rect.top + window.scrollY;
    shellHeight = rect.height;
  }

  function onScroll() {
    const y = window.scrollY;

    if (topbar) {
      if (y > 90 && y > lastY + 2) {
        topbar.classList.add('topbar-hidden');
      } else {
        topbar.classList.remove('topbar-hidden');
      }
    }

    if (shell) {
      if (y >= shellTop - 6) {
        if (!shell.classList.contains('is-docked')) {
          shell.classList.add('is-docked');
          spacer.classList.add('active');
          spacer.style.height = `${shellHeight}px`;
        }
      } else {
        shell.classList.remove('is-docked');
        spacer.classList.remove('active');
        spacer.style.height = '0px';
      }
    }

    lastY = y;
  }

  recalc();
  onScroll();
  window.addEventListener('scroll', onScroll, { passive: true });
  window.addEventListener('resize', () => {
    recalc();
    onScroll();
  });
}

initDockedSearchAndNavbar();

function formatIDR(num) {
  return `IDR ${Number(num).toLocaleString('id-ID')}`;
}

function initHotelFilterPanel() {
  const grid = document.getElementById('hotelGrid');
  if (!grid) return;

  const cards = Array.from(grid.querySelectorAll('.hotel-card'));
  cards.forEach((card, idx) => {
    const stars = [3, 4, 5][idx % 3];
    const guest = [7.4, 8.2, 9.1][idx % 3];
    const promo = idx % 2 === 0 ? 'promo' : 'ramadan';
    const baseFacilities = ['wifi', 'pool', 'breakfast', 'gym', 'airport'];
    card.dataset.star = String(stars);
    card.dataset.guest = String(guest);
    card.dataset.promo = promo;
    card.dataset.facilities = baseFacilities.filter((_, i) => (idx + i) % 2 === 0).join(',');
    card.dataset.price = String(450000 + (idx % 10) * 150000);
  });

  const maxPriceInput = document.getElementById('maxPrice');
  const maxPriceText = document.getElementById('maxPriceText');
  const empty = document.getElementById('emptyHotelResult');
  const filters = Array.from(document.querySelectorAll('.hotel-filter'));

  function applyFilters() {
    const selected = {
      facility: filters.filter((f) => f.checked && f.dataset.type === 'facility').map((f) => f.value),
      star: filters.filter((f) => f.checked && f.dataset.type === 'star').map((f) => f.value),
      guest: filters.filter((f) => f.checked && f.dataset.type === 'guest').map((f) => Number(f.value)),
      promo: filters.filter((f) => f.checked && f.dataset.type === 'promo').map((f) => f.value),
    };

    const maxPrice = Number(maxPriceInput?.value || 2000000);
    if (maxPriceText) maxPriceText.innerText = formatIDR(maxPrice);

    let visible = 0;
    cards.forEach((card) => {
      const cardFacilities = (card.dataset.facilities || '').split(',').filter(Boolean);
      const cardStar = card.dataset.star || '0';
      const cardGuest = Number(card.dataset.guest || 0);
      const cardPromo = card.dataset.promo || '';
      const cardPrice = Number(card.dataset.price || 0);

      const facilityPass = selected.facility.length === 0 || selected.facility.every((f) => cardFacilities.includes(f));
      const starPass = selected.star.length === 0 || selected.star.includes(cardStar);
      const guestPass = selected.guest.length === 0 || selected.guest.some((g) => cardGuest >= g);
      const promoPass = selected.promo.length === 0 || selected.promo.includes(cardPromo);
      const pricePass = cardPrice <= maxPrice;

      const show = facilityPass && starPass && guestPass && promoPass && pricePass;
      card.style.display = show ? '' : 'none';
      if (show) visible += 1;
    });

    if (empty) empty.classList.toggle('hidden', visible > 0);
  }

  document.getElementById('applyAllFilters')?.addEventListener('click', applyFilters);
  document.getElementById('clearAllFilters')?.addEventListener('click', () => {
    filters.forEach((f) => { f.checked = false; });
    if (maxPriceInput) maxPriceInput.value = '2000000';
    applyFilters();
  });
  document.getElementById('resetFilters')?.addEventListener('click', () => {
    filters.forEach((f) => { f.checked = false; });
    applyFilters();
  });
  document.getElementById('resetPrice')?.addEventListener('click', () => {
    if (maxPriceInput) maxPriceInput.value = '2000000';
    applyFilters();
  });

  maxPriceInput?.addEventListener('input', applyFilters);
  filters.forEach((f) => f.addEventListener('change', applyFilters));

  document.querySelectorAll('[data-accordion] .acc-head').forEach((head) => {
    head.addEventListener('click', () => {
      head.parentElement.classList.toggle('collapsed');
    });
  });

  document.querySelectorAll('#detailTabs a').forEach((a) => {
    a.addEventListener('click', (e) => {
      const id = a.getAttribute('href');
      const target = id ? document.querySelector(id) : null;
      if (!target) return;
      e.preventDefault();
      target.scrollIntoView({ behavior: 'smooth', block: 'start' });
    });
  });

  applyFilters();
}

initHotelFilterPanel();
