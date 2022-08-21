package leetcode
/*
标题 : Design Authentication Manager
链接 : https://leetcode.cn/problems/design-authentication-manager/
难度 : 一般
解题思路:


注意事项:

*/

type AuthenticationManager struct {

}


func Constructor(timeToLive int) AuthenticationManager {

}


func (this *AuthenticationManager) Generate(tokenId string, currentTime int)  {

}


func (this *AuthenticationManager) Renew(tokenId string, currentTime int)  {

}


func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {

}


/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */