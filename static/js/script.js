function openPopup(title,img,desc,id){
  const popup=document.getElementById('popup');
  document.getElementById('ptitle').innerText=title;
  document.getElementById('pimg').src=img;
  document.getElementById('pdesc').innerText=desc;
  document.getElementById('pdetail').href='/hotel?id='+id;
  popup.classList.remove('hidden');
}

function closePopup(){
  document.getElementById('popup').classList.add('hidden');
}

document.addEventListener('click',(e)=>{
  const popup=document.getElementById('popup');
  if(!popup.classList.contains('hidden') && e.target.id==='popup') closePopup();
});

document.addEventListener('keydown',(e)=>{
  if(e.key==='Escape') closePopup();
});
