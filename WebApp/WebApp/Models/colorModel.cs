using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;

namespace WebApp.Models
{
    public class colorModel
    {
        [BsonElement("color")]
        public string color { get; set; }
    }
}
