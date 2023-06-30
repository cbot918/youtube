const cloudinary = require('cloudinary')

cloudinary.config({ 
  cloud_name: 'yalecloud', 
  api_key: '119372989697243', 
  api_secret: 'tfwtQFL_JGO5_J70_0ABJkpxGHo' 
});

cloudinary.v2.uploader.upload("./bighead.png",
  { public_id: "olympic_flag" }, 
  function(error, result) {console.log(result); });