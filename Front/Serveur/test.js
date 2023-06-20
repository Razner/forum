var likeButton = document.getElementById('like-button');
var dislikeButton = document.getElementById('dislike-button');

var likeImage = document.getElementById('like-image');
var dislikeImage = document.getElementById('dislike-image');

var isLiked = false;
var isDisliked = false;

likeButton.addEventListener('click', function() {
  if (isLiked) {
    likeImage.src = '../assets/images/LikeWhite.png';
    isLiked = false;
  } else {
    likeImage.src = '../assets/images/LikeGreen.png';
    isLiked = true;
  }
});

dislikeButton.addEventListener('click', function() {
  if (isDisliked) {
    dislikeImage.src = '../assets/images/DislikeWhite.png';
    isDisliked = false;
  } else {
    dislikeImage.src = '../assets/images/DislikeRed.png';
    isDisliked = true;
  }
});
