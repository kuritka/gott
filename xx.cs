using System;
using System.Collections;
using System.Globalization;
using System.Reflection;

public class Example : IComparer
{
   public static void Main()
   {
   	  var cultureCode = "sq-AL";
   	  var culture = CultureInfo.GetCultureInfo(cultureCode);

      NumberFormatInfo nfi1 = culture.NumberFormat;

      PropertyInfo[] props = nfi1.GetType().GetProperties(BindingFlags.Instance | BindingFlags.Public);
      Array.Sort(props, new Example());
      Console.WriteLine("Properties of NumberFormat.CurrentInfo object:");
      foreach (var prop in props) {
         if (prop.PropertyType.IsArray) {
            Array arr = prop.GetValue(nfi1) as Array;
            Console.Write(String.Format("   {0}: ", prop.Name) + "{ ");
            int ctr = 0;
            foreach (var item in arr) {
               Console.Write("{0}{1}", item, ctr == arr.Length - 1 ?" }" : ", ");
               ctr++;
            }
            Console.WriteLine();
         }
         else {
            Console.WriteLine("   {0}: {1}", prop.Name, prop.GetValue(nfi1));
        }
      }

	  Console.WriteLine("");
	  Console.Write(String.Format(" \"{0}\": ", cultureCode) + "{");
	 foreach (var prop in props) {

	 	if (prop.Name == "DigitSubstitution" || prop.Name == "IsReadOnly" || prop.Name == "NativeDigits" || prop.Name == "NaNSymbol")
			continue;

         if (prop.PropertyType.IsArray) {
            Array arr = prop.GetValue(nfi1) as Array;
            Console.Write(String.Format("{0}: ", prop.Name) + "[]int32{ ");
            int ctr = 0;
            foreach (var item in arr) {
               Console.Write("{0}{1}", item, ctr == arr.Length - 1 ?" }" : ", ");
               ctr++;
            }
            Console.Write(",");
         }
         else {
		 if(prop.GetValue(nfi1).GetType() == typeof(string)){
            Console.Write("{0}: \"{1}\",", prop.Name, prop.GetValue(nfi1));
		 } else {
		 	Console.Write("{0}: {1},", prop.Name, prop.GetValue(nfi1));
		 }
        }
      }
	  Console.Write("},");

   }

   public int Compare(Object x, Object y)
   {
      if (x == null && y == null) return 0;

      PropertyInfo px = x as PropertyInfo;
      if (px == null) return -1;

      PropertyInfo py = y as PropertyInfo;
      if (py == null) return 1;

      return String.Compare(px.Name, py.Name);
   }
}
// The example displays the following output:
//       Objects equal: True
//       Equal references: True
//
//       Properties of NumberFormat.CurrentInfo object:
//          CurrencyDecimalDigits: 2
//          CurrencyDecimalSeparator: .
//          CurrencyGroupSeparator: ,
//          CurrencyGroupSizes: { 3 }
//          CurrencyNegativePattern: 0
//          CurrencyPositivePattern: 0
//          CurrencySymbol: $
//          DigitSubstitution: None
//          IsReadOnly: True
//          NaNSymbol: NaN
//          NativeDigits: { 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }
//          NegativeInfinitySymbol: -Infinity
//          NegativeSign: -
//          NumberDecimalDigits: 2
//          NumberDecimalSeparator: .
//          NumberGroupSeparator: ,
//          NumberGroupSizes: { 3 }
//          NumberNegativePattern: 1
//          PercentDecimalDigits: 2
//          PercentDecimalSeparator: .
//          PercentGroupSeparator: ,
//          PercentGroupSizes: { 3 }
//          PercentNegativePattern: 0
//          PercentPositivePattern: 0
//          PercentSymbol: %
//          PerMilleSymbol: %
//          PositiveInfinitySymbol: Infinity
//          PositiveSign: +