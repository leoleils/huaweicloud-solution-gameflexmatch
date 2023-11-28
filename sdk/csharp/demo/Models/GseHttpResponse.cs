using System;

namespace CSharpDemo.Models
{
    public class GseHttpResponse
    {   
        public int code;
        public string message;
        public object result;

        public GseHttpResponse()
        {
            code = 0;
            message = "SUCCESS";
            result = new object{};
        }

        public GseHttpResponse(int _code, string _message, object _result)
        {
            code = _code;
            message = _message;
            if(_result == null)
            {
                result = new object{};
            }
            else
            {
                result = _result;
            }
            result = _result;
        }
    }
}