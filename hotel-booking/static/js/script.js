(() => {
  const cards = document.querySelectorAll('.city-card');
  cards.forEach(c => c.addEventListener('mouseenter', () => { c.style.cursor = 'pointer'; }));

  const searchInput = document.getElementById('searchInput');
  const results = document.querySelectorAll('.result-item');
  const cityCards = document.querySelectorAll('.city-card');

  if (searchInput) {
    searchInput.addEventListener('input', () => {
      const key = searchInput.value.trim().toLowerCase();
      let visibleResult = 0;

      results.forEach((item) => {
        const hotel = (item.dataset.hotel || '').toLowerCase();
        const city = (item.dataset.city || '').toLowerCase();
        const show = key === '' || hotel.includes(key) || city.includes(key);
        item.classList.toggle('hidden', !show);
        if (show) visibleResult++;
      });

      cityCards.forEach((card) => {
        const city = (card.dataset.city || '').toLowerCase();
        card.style.display = key === '' || city.includes(key) ? '' : 'none';
      });

      if (key !== '' && visibleResult === 0) {
        const panel = document.getElementById('searchResults');
        if (panel && !document.getElementById('emptyResult')) {
          const el = document.createElement('div');
          el.id = 'emptyResult';
          el.className = 'result-item';
          el.innerHTML = '<span>Tidak ada kota/hotel yang cocok.</span>';
          panel.appendChild(el);
        }
      } else {
        const empty = document.getElementById('emptyResult');
        if (empty) empty.remove();
      }
    });
  }

  const guestTrigger = document.getElementById('guestTrigger');
  const guestPopover = document.getElementById('guestPopover');
  const summary = document.getElementById('guestSummary');
  const counts = {
    adult: document.getElementById('adultCount'),
    child: document.getElementById('childCount'),
    room: document.getElementById('roomCount')
  };

  const syncSummary = () => {
    if (!summary) return;
    summary.textContent = `${counts.adult.textContent} Dewasa, ${counts.child.textContent} Anak, ${counts.room.textContent} Kamar`;
  };

  if (guestTrigger && guestPopover) {
    guestTrigger.addEventListener('click', () => guestPopover.classList.toggle('open'));

    guestPopover.querySelectorAll('button[data-target]').forEach((btn) => {
      btn.addEventListener('click', () => {
        const target = btn.dataset.target;
        const act = btn.dataset.act;
        const node = counts[target];
        if (!node) return;
        const min = target === 'adult' || target === 'room' ? 1 : 0;
        const max = target === 'adult' ? 10 : 6;
        let value = Number(node.textContent);
        value = act === 'inc' ? value + 1 : value - 1;
        value = Math.min(max, Math.max(min, value));
        node.textContent = String(value);
        syncSummary();
      });
    });

    document.addEventListener('click', (event) => {
      if (!guestPopover.contains(event.target) && !guestTrigger.contains(event.target)) {
        guestPopover.classList.remove('open');
      }
    });
  }
})();
