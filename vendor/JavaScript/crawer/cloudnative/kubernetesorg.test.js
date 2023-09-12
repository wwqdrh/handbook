const { GetNewArticle, GetArticleContent } = require('./kubernetesorg');

test('getNewArticle', () => {
  GetNewArticle();
});

test('getArticleContent', () => {
  GetArticleContent("https://www.kubernetes.org.cn/9843.html");
})